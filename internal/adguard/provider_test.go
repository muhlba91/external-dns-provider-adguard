package adguard

import (
	"context"
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type testCase struct {
	name           string
	filteringRules getFilteringRules
	hasError       bool
	requestError   error
	endpoints      []*endpoint.Endpoint
	changes        *plan.Changes
	rules          []string
	domainFilter   endpoint.DomainFilter
	log.Ext1FieldLogger
}

var mockHTTPClient *MockHTTPClient
var testProvider *Provider

func TestNewAdguardProvider(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	testCases := []struct {
		name     string
		config   *Configuration
		hasError bool
	}{
		{
			name: "minimal provider config",
			config: &Configuration{
				URL:    "https://domain.com",
				DryRun: true,
			},
		},
		{
			name:     "errornous provider config",
			hasError: true,
			config: &Configuration{
				URL:    "my domain",
				DryRun: false,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d. %s", i+1, tc.name), func(t *testing.T) {
			p, err := NewAdguardProvider(endpoint.DomainFilter{}, tc.config)
			if tc.hasError {
				require.Nil(t, p)
				require.Error(t, err)
			} else {
				require.NotNil(t, p)
				require.NoError(t, err)
			}
		})
	}
}

func TestEndpointSupported(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	testCases := []struct {
		name     string
		endpoint *endpoint.Endpoint
		valid    bool
	}{
		{
			name:     "A record",
			endpoint: endpoint.NewEndpoint("domain.com", endpoint.RecordTypeA, "1.1.1.1"),
			valid:    true,
		},
		{
			name:     "AAAA record",
			endpoint: endpoint.NewEndpoint("domain.com", endpoint.RecordTypeAAAA, "1111:1111::1"),
			valid:    true,
		},
		{
			name:     "TXT record",
			endpoint: endpoint.NewEndpoint("domain.com", endpoint.RecordTypeTXT, "text"),
			valid:    true,
		},
		{
			name:     "CNAME record",
			endpoint: endpoint.NewEndpoint("domain.com", endpoint.RecordTypeCNAME, "other.org"),
			valid:    true,
		},
		{
			name:     "SRV record",
			endpoint: endpoint.NewEndpoint("domain.com", endpoint.RecordTypeSRV, "rsv"),
			valid:    false,
		},
		{
			name:     "NS record",
			endpoint: endpoint.NewEndpoint("domain.com", endpoint.RecordTypeNS, "1.1.1.1"),
			valid:    false,
		},
		{
			name:     "PTR record",
			endpoint: endpoint.NewEndpoint("1.1.1.1", endpoint.RecordTypePTR, "domain.com"),
			valid:    false,
		},
		{
			name:     "MX record",
			endpoint: endpoint.NewEndpoint("1.1.1.1", endpoint.RecordTypeMX, "10 mail.domain.com."),
			valid:    false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d. %s", i+1, tc.name), func(t *testing.T) {
			require.Equal(t, tc.valid, endpointSupported(tc.endpoint))
		})
	}
}

func TestDeserializeToEndpoint(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	testCases := []struct {
		name        string
		text        string
		endpoint    *endpoint.Endpoint
		expectedErr bool
	}{
		{
			name:     "A record",
			text:     fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeA, Targets: []string{"1.1.1.1"}},
		},
		{
			name:     "AAAA record",
			text:     fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy),
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeAAAA, Targets: []string{"1111:1111::1"}},
		},
		{
			name:     "TXT record",
			text:     fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;TXT;external-dns-txt #%s", managedBy),
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeTXT, Targets: []string{"external-dns-txt"}},
		},
		{
			name:     "CNAME record",
			text:     fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;CNAME;other.org #%s", managedBy),
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeCNAME, Targets: []string{"other.org"}},
		},
		{
			name:        "invalid record",
			text:        fmt.Sprintf("@@||abc.com #%s", managedBy),
			expectedErr: true,
		},
		{
			name:        "unmanaged record",
			text:        "@@||abc.com",
			expectedErr: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d. %s", i+1, tc.name), func(t *testing.T) {
			ep, err := deserializeToEndpoint(tc.text)
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.endpoint, ep)
			}
		})
	}
}

func TestSerializeToString(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	testCases := []struct {
		name     string
		text     []string
		endpoint *endpoint.Endpoint
	}{
		{
			name:     "A record",
			text:     []string{fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy)},
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeA, Targets: []string{"1.1.1.1"}},
		},
		{
			name:     "AAAA record",
			text:     []string{fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy)},
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeAAAA, Targets: []string{"1111:1111::1"}},
		},
		{
			name:     "TXT record",
			text:     []string{fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;TXT;external-dns-txt #%s", managedBy)},
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeTXT, Targets: []string{"external-dns-txt"}},
		},
		{
			name:     "CNAME record",
			text:     []string{fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;CNAME;other.org #%s", managedBy)},
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeCNAME, Targets: []string{"other.org"}},
		},
		{
			name: "multiple records",
			text: []string{
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
			},
			endpoint: &endpoint.Endpoint{DNSName: "domain.com", RecordType: endpoint.RecordTypeA, Targets: []string{"1.1.1.1", "2.2.2.2"}},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d. %s", i+1, tc.name), func(t *testing.T) {
			rr := serializeToString(tc.endpoint)
			require.Equal(t, tc.text, rr)
		})
	}
}

func TestRecords(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	testCases := []*testCase{
		{
			name:         "valid case",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy),
					fmt.Sprintf("||other.org^$dnsrewrite=NOERROR;A;3.3.3.3 #%s", managedBy),
				},
			},
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "domain.com",
					RecordType: endpoint.RecordTypeA,
					Targets: []string{
						"1.1.1.1",
						"2.2.2.2",
					},
				},
				{
					DNSName:    "domain.com",
					RecordType: endpoint.RecordTypeAAAA,
					Targets: []string{
						"1111:1111::1",
					},
				},
				{
					DNSName:    "other.org",
					RecordType: endpoint.RecordTypeA,
					Targets: []string{
						"3.3.3.3",
					},
				},
			},
		},
		{
			name:         "unmanaged filters",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy),
					"||other.org^$dnsrewrite=NOERROR;A;3.3.3.3 #unmanaged",
					"@@||other.org",
				},
			},
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "domain.com",
					RecordType: endpoint.RecordTypeA,
					Targets: []string{
						"1.1.1.1",
						"2.2.2.2",
					},
				},
				{
					DNSName:    "domain.com",
					RecordType: endpoint.RecordTypeAAAA,
					Targets: []string{
						"1111:1111::1",
					},
				},
			},
		},
		{
			name:         "valid case with domain filter",
			hasError:     false,
			domainFilter: endpoint.NewDomainFilter([]string{"domain.com"}),
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy),
					fmt.Sprintf("||other.org^$dnsrewrite=NOERROR;A;3.3.3.3 #%s", managedBy),
				},
			},
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "domain.com",
					RecordType: endpoint.RecordTypeA,
					Targets: []string{
						"1.1.1.1",
						"2.2.2.2",
					},
				},
				{
					DNSName:    "domain.com",
					RecordType: endpoint.RecordTypeAAAA,
					Targets: []string{
						"1111:1111::1",
					},
				},
			},
		},
		{
			name:         "invalid filters",
			hasError:     true,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 whatever #%s", managedBy),
				},
			},
		},
		{
			name:         "request error",
			hasError:     true,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{},
			},
			requestError: fmt.Errorf("error"),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d. %s", i+1, tc.name), func(t *testing.T) {
			mockHTTPClient = &MockHTTPClient{
				testCase: tc,
				t:        t,
			}
			testProvider = &Provider{
				client:       mockHTTPClient,
				domainFilter: tc.domainFilter,
			}

			records, err := testProvider.Records(context.TODO())
			if tc.hasError {
				require.Error(t, err)
			} else {
				require.ElementsMatch(t, tc.endpoints, records)
			}
		})
	}
}

func TestApplyChanges(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	testCases := []*testCase{
		{
			name:         "valid create",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			rules: []string{
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy),
				fmt.Sprintf("||other.org^$dnsrewrite=NOERROR;A;3.3.3.3 #%s", managedBy),
			},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"1.1.1.1",
							"2.2.2.2",
						},
					},
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeAAAA,
						Targets: []string{
							"1111:1111::1",
						},
					},
					{
						DNSName:    "other.org",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"3.3.3.3",
						},
					},
				},
			},
		},
		{
			name:         "valid update",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			rules: []string{
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
			},
			changes: &plan.Changes{
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"2.2.2.2",
						},
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"1.1.1.1",
						},
					},
				},
			},
		},
		{
			name:         "valid delete",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
				},
			},
			rules: []string{},
			changes: &plan.Changes{
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"2.2.2.2",
						},
					},
				},
			},
		},
		{
			name:         "valid partial delete",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
				},
			},
			rules: []string{
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
			},
			changes: &plan.Changes{
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"2.2.2.2",
						},
					},
				},
			},
		},
		{
			name:         "valid delete and create",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
				},
			},
			rules: []string{
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy),
			},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeAAAA,
						Targets: []string{
							"1111:1111::1",
						},
					},
				},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"2.2.2.2",
						},
					},
				},
			},
		},
		{
			name:         "valid create and update",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			rules: []string{
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;1.1.1.1 #%s", managedBy),
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;3.3.3.3 #%s", managedBy),
			},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"3.3.3.3",
						},
					},
				},
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"2.2.2.2",
						},
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeA,
						Targets: []string{
							"1.1.1.1",
						},
					},
				},
			},
		},
		{
			name:         "valid unmanaged rules filter",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
					"||other.org^$dnsrewrite=NOERROR;A;3.3.3.3 #unmanaged",
					"@@||other.org",
				},
			},
			rules: []string{
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 #%s", managedBy),
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy),
				"||other.org^$dnsrewrite=NOERROR;A;3.3.3.3 #unmanaged",
				"@@||other.org",
			},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeAAAA,
						Targets: []string{
							"1111:1111::1",
						},
					},
				},
			},
		},
		{
			name:         "invalid type",
			hasError:     false,
			domainFilter: endpoint.DomainFilter{},
			rules: []string{
				fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;AAAA;1111:1111::1 #%s", managedBy),
			},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeSRV,
						Targets: []string{
							"srv",
						},
					},
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeAAAA,
						Targets: []string{
							"1111:1111::1",
						},
					},
				},
			},
		},
		{
			name:         "invalid existing rule",
			hasError:     true,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{
					fmt.Sprintf("||domain.com^$dnsrewrite=NOERROR;A;2.2.2.2 whatever #%s", managedBy),
				},
			},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "domain.com",
						RecordType: endpoint.RecordTypeAAAA,
						Targets: []string{
							"1111:1111::1",
						},
					},
				},
			},
		},
		{
			name:         "request error",
			hasError:     true,
			domainFilter: endpoint.DomainFilter{},
			filteringRules: getFilteringRules{
				UserRules: []string{},
			},
			requestError: fmt.Errorf("error"),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d. %s", i+1, tc.name), func(t *testing.T) {
			mockHTTPClient = &MockHTTPClient{
				testCase: tc,
				t:        t,
			}
			testProvider = &Provider{
				client:       mockHTTPClient,
				domainFilter: tc.domainFilter,
			}

			err := testProvider.ApplyChanges(context.TODO(), tc.changes)
			if tc.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

type MockHTTPClient struct {
	t        *testing.T
	testCase *testCase
}

func (d *MockHTTPClient) GetFilteringRules(_ context.Context) ([]string, error) {
	return d.testCase.filteringRules.UserRules, d.testCase.requestError
}

func (d *MockHTTPClient) SetFilteringRules(_ context.Context, rules []string) error {
	require.ElementsMatch(d.t, d.testCase.rules, rules)
	return nil
}
