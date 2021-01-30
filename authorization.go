package currencycloud_go

import (
	"bytes"
	"context"
	"net/http"
)

// Login is the Currencycloud API authentication and authorization endpoint
func (c *Client) Login(ctx context.Context) (*AuthTokenResponse, error) {
	cred := c.GetCredentials().Encode()
	buf := bytes.NewBuffer([]byte(cred))

	req, err := http.NewRequestWithContext(ctx, "POST", c.applyApiBaseUrl(EndpointLogin), buf)
	if err != nil {
		return &AuthTokenResponse{}, err
	}

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	response := &AuthTokenResponse{}

	err = c.Send(req, response)

	if response.AuthToken != "" {
		c.authToken = response
	}

	return response, err
}

// Logout to retire its authentication token early rather
func (c *Client) Logout(ctx context.Context) error {
	buf := bytes.NewBuffer([]byte(""))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.applyApiBaseUrl(EndpointLogin), buf)

	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
