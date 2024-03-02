package adguard

import (
	"context"
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// Provider type for interfacing with Adguard
type Provider struct {
	provider.BaseProvider

	client       Client
	domainFilter endpoint.DomainFilter
}

var (
	errNotManaged = fmt.Errorf("not managed by external-dns")
)

// NewAdguardProvider initializes a new provider
func NewAdguardProvider(domainFilter endpoint.DomainFilter, config *Configuration) (provider.Provider, error) {
	log.Debugf("using adguard at %s", config.URL)

	// URL adjustment according to the specification
	if !strings.HasSuffix(config.URL, "/") {
		config.URL = config.URL + "/"
	}
	if !strings.HasSuffix(config.URL, "control/") {
		config.URL = config.URL + "control/"
	}

	c, err := newAdguardClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create the adguard client: %w", err)
	}

	p := &Provider{
		client:       c,
		domainFilter: domainFilter,
	}

	return p, nil
}

// ApplyChanges syncs the desired state with Adguard
func (p *Provider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Debugf("received changes: %+v", changes)

	or, err := p.client.GetFilteringRules(ctx)
	if err != nil {
		return err
	}

	// resulting rules
	rr := make([]string, 0)
	// list of endpoints to create - endpoints are referenced by dns name and record type
	eps := make(map[string]*endpoint.Endpoint)

	// collect all non-managed rules and existing endpoints managed by external-dns
	for _, r := range or {
		ep, err := deserializeToEndpoint(r)
		if err != nil {
			// rules not managed by external-dns are kept
			if errors.Is(err, errNotManaged) {
				rr = append(rr, r)
				continue
			}
			return fmt.Errorf("failed to parse rule %s: %w", r, err)
		}

		epk := ep.DNSName + ep.RecordType
		if eps[epk] != nil {
			eps[epk].Targets = append(eps[epk].Targets, ep.Targets...)
		} else {
			eps[epk] = ep
		}
	}

	// delete all records to be updated or deleted
	for _, dep := range append(changes.UpdateOld, changes.Delete...) {
		epk := dep.DNSName + dep.RecordType
		if ep, ok := eps[epk]; ok {
			for _, t := range dep.Targets {
				if slices.Contains(ep.Targets, t) {
					ti := slices.Index(ep.Targets, t)
					ep.Targets = append(ep.Targets[:ti], ep.Targets[ti+1:]...)
					log.Debugf("deleting target %s for %s %s", t, dep.DNSName, dep.RecordType)
				}
				if len(ep.Targets) == 0 {
					delete(eps, epk)
					log.Debugf("deleting rule %s %s", dep.DNSName, dep.RecordType)
					break
				}
			}
		}
	}

	// add all endpoints and targets to be created
	for _, cep := range append(changes.Create, changes.UpdateNew...) {
		if !endpointSupported(cep) {
			log.Warnf("requested unsupported endpoint creation: %s", cep)
			continue
		}

		epk := cep.DNSName + cep.RecordType
		if ep, ok := eps[epk]; ok {
			ep.Targets = append(ep.Targets, cep.Targets...)
			log.Debugf("adding target %s to existing rule for %s %s", ep.Targets, ep.DNSName, ep.RecordType)
		} else {
			ep = &endpoint.Endpoint{
				DNSName:    cep.DNSName,
				RecordType: cep.RecordType,
				Targets:    cep.Targets,
			}
			eps[epk] = ep
			log.Debugf("adding rule %s", cep)
		}
	}

	// convert endpoints to rules
	for _, e := range eps {
		s := serializeToString(e)
		rr = append(rr, s...)
	}

	return p.client.SetFilteringRules(ctx, rr)
}

// Records reads all endpoints from Adguard
func (p *Provider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	resp, err := p.client.GetFilteringRules(ctx)
	if err != nil {
		return nil, err
	}

	// deserialize all endpoints managed by external-dns
	// endpoints are referenced by dns name and record type
	eps := make(map[string]*endpoint.Endpoint)
	for _, rule := range resp {
		ep, err := deserializeToEndpoint(rule)
		if err != nil {
			// unmanaged rules are ignored
			if err == errNotManaged {
				continue
			}
			return nil, err
		}

		if !p.domainFilter.Match(ep.DNSName) {
			continue
		}

		epk := ep.DNSName + ep.RecordType
		if eep, ok := eps[epk]; ok {
			eep.Targets = append(eep.Targets, ep.Targets...)
			log.Debugf("found target %s for existing rule for %s %s", ep.Targets, ep.DNSName, ep.RecordType)
		} else {
			eps[epk] = ep
			log.Debugf("found rule %s", ep)
		}
	}

	return maps.Values(eps), nil
}

func endpointSupported(e *endpoint.Endpoint) bool {
	// Adguard does not have any restriction, and we can allow all upstream/external-dns ones
	return e.RecordType == endpoint.RecordTypeA ||
		e.RecordType == endpoint.RecordTypeTXT ||
		e.RecordType == endpoint.RecordTypeAAAA ||
		e.RecordType == endpoint.RecordTypeCNAME ||
		e.RecordType == endpoint.RecordTypeSRV ||
		e.RecordType == endpoint.RecordTypeNS ||
		e.RecordType == endpoint.RecordTypePTR ||
		e.RecordType == endpoint.RecordTypeMX
}

func deserializeToEndpoint(rule string) (*endpoint.Endpoint, error) {
	// format: "|DNS.NAME^dnsrewrite=NOERROR;RECORD_TYPE;TARGET"
	p := strings.SplitN(rule, ";", 3)
	if len(p) != 3 {
		return nil, errNotManaged
	}
	dp := strings.SplitN(p[0], "^", 2)
	if strings.HasPrefix(dp[0], "||") {
		return nil, errNotManaged
	}
	if len(dp) != 2 {
		return nil, fmt.Errorf("invalid rule: %s", rule)
	}
	d := strings.TrimPrefix(dp[0], "|")

	// see serializeToString for the format
	r := &endpoint.Endpoint{
		RecordType: p[1],
		DNSName:    d,
		Targets:    endpoint.Targets{p[2]},
	}

	return r, nil
}

func serializeToString(e *endpoint.Endpoint) []string {
	r := []string{}
	for _, t := range e.Targets {
		// format: "|DNS.NAME^dnsrewrite=NOERROR;RECORD_TYPE;TARGET"
		r = append(r, fmt.Sprintf("|%s^$dnsrewrite=NOERROR;%s;%s", e.DNSName, e.RecordType, t))
	}
	return r
}
