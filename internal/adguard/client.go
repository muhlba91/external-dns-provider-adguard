package adguard

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Client interface for interacting with Adguard
// See the OpenAPI spec: https://raw.githubusercontent.com/AdguardTeam/AdGuardHome/master/openapi/openapi.yaml
type Client interface {
	GetFilteringRules(ctx context.Context) ([]string, error)
	SetFilteringRules(ctx context.Context, rules []string) error
	Status(ctx context.Context) (bool, error)
}

// hhtpClient type used to implement Client with an HTTP client
type httpClient struct {
	hc     *http.Client
	config *Configuration
}

// adguardStatus is the response of retrieving the AdGuard status
type adguardStatus struct {
	Version string `json:"version"`
	Running bool   `json:"running"`
}

// getFilteringRules is the response of retrieving filtering rules
type getFilteringRules struct {
	UserRules []string `json:"user_rules"`
}

// setFilteringRules is the request sent to Adguard for setting new rules
type setFilteringRules struct {
	Rules []string `json:"rules"`
}

// newAdguardClient initializes a new HTTP client
func newAdguardClient(config *Configuration) (*httpClient, error) {
	hc := http.Client{}
	c := &httpClient{
		hc:     &hc,
		config: config,
	}

	// check validity of the configuration
	s, err := c.Status(context.Background())
	if err != nil || !s {
		return nil, err
	}

	return c, nil
}

func (c *httpClient) doRequest(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	log.Debugf("making %s request to /%s", method, path)

	req, err := http.NewRequestWithContext(ctx, method, c.config.URL+path, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.config.User, c.config.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	log.Debugf("response code from %s request to %s: %d", method, path, resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s request to %s was not successful: %d", method, path, resp.StatusCode)
	}

	return resp, nil
}

func (c *httpClient) Status(ctx context.Context) (bool, error) {
	if c.config.DryRun {
		log.Info("would check adguard configuration")
		return true, nil
	}

	r, err := c.doRequest(ctx, http.MethodGet, "status", nil)
	if err != nil {
		return false, err
	}
	defer r.Body.Close()

	var resp adguardStatus
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		return false, err
	}
	log.Debugf("retrieved status: %+v", resp)

	return resp.Running, nil
}

// Retrieves all existing filtering rules from Adguard
func (c *httpClient) GetFilteringRules(ctx context.Context) ([]string, error) {
	if c.config.DryRun {
		log.Info("would retrieve rules")
		return []string{}, nil
	}

	r, err := c.doRequest(ctx, http.MethodGet, "filtering/status", nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var resp getFilteringRules
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	log.Debugf("retrieved filtering rules: %+v", resp)

	return resp.UserRules, nil
}

// Sets new filtering rules in Adguard
func (c *httpClient) SetFilteringRules(ctx context.Context, rules []string) error {
	if c.config.DryRun {
		log.Infof("would set rules: %+v", rules)
		return nil
	}

	body := setFilteringRules{Rules: rules}
	log.Debugf("sending filtering rules: %s", body)

	b := bytes.NewBuffer(nil)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return err
	}

	r, err := c.doRequest(ctx, http.MethodPost, "filtering/set_rules", b)
	if err != nil {
		return err
	}
	_ = r.Body.Close()

	return nil
}
