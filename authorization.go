package currencycloud_go

import (
	"bytes"
	"context"
	"net/http"
)

// Login is the Currencycloud API authentication and authorization endpoint
func (c *Client) Login(ctx context.Context) (*AuthTokenResponse, error) {
	resp := &AuthTokenResponse{}

	buf := bytes.NewBuffer([]byte(c.GetCredentials().Encode()))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.applyApiBaseUrl(EndpointLogin), buf)
	if err != nil {
		return resp, err
	}
	req.Header.Set("Content-type", HeaderFormUrlencoded)

	if err = c.Send(req, resp); err != nil {
		return resp, err
	}

	c.SetAuthToken(resp.AuthToken)

	return resp, nil
}

// Logout to retire its authentication token early rather
func (c *Client) Logout(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.applyApiBaseUrl(EndpointLogout), nil)
	if err != nil {
		return err
	}
	return c.SendWithAuth(req, nil)
}
