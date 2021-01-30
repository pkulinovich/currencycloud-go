package currencycloud_go

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// NewClient returns new Client struct
func NewClient(loginID string, apiKey string, env string) (*Client, error) {
	if loginID == "" || apiKey == "" || env == "" {
		return nil, errors.New("LoginID, ApiKey and Env are required to create a Client")
	}

	return &Client{
		client:  &http.Client{},
		loginID: loginID,
		apiKey:  apiKey,
		env:     Environment(env),
	}, nil
}

// NewRequest constructs a request
func (c *Client) NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequestWithContext(ctx, method, url, buf)
}

// Send makes a request to the API.
func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", HeaderContentTypeApplicationJson)
	req.Header.Set("User-Agent", HeaderUserAgent)

	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", HeaderContentTypeApplicationJson)
	}

	resp, err = c.client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err = ioutil.ReadAll(resp.Body)
		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}
		return errResp
	}
	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		io.Copy(w, resp.Body)
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// SendWithAuth makes a request to the API and apply X-Auth-Token header automatically
func (c *Client) SendWithAuth(req *http.Request, v interface{}) error {
	c.Lock() // we don't need to `defer c.Unlock()` because we need `c.Send(...)`

	if c.authToken != "" {
		if !c.tokenExpiresAt.IsZero() && c.tokenExpiresAt.Sub(time.Now()) < ExpiresInLimit {
			if _, err := c.Login(req.Context()); err != nil {
				c.Unlock()
				return err
			}
		}
		req.Header.Set("X-Auth-Token", c.authToken)
	}

	c.Unlock() // Unlock the client mutex before sending the request
	return c.Send(req, v)
}

// SetAuthToken sets saved token to current client
func (c *Client) SetAuthToken(token string) {
	c.authToken = token

	// Tokens expire after 30 minutes of inactivity.
	c.tokenExpiresAt = time.Now().Add(time.Minute * 30)

	// TODO: Token requests are limited to 10 calls per minute.
}

// SetLog will set/change the output destination.
func (c *Client) SetLogger(logger io.Writer) {
	c.logger = logger
}

// GetCredentials retrieve credentials
func (c *Client) GetCredentials() url.Values {
	v := url.Values{}
	v.Set("login_id", c.loginID)
	v.Set("api_key", c.apiKey)
	return v
}

func (c *Client) applyApiBaseUrl(path Endpoint) string {
	return fmt.Sprintf("%s%s", urls[c.env], path)
}

// log will dump request and response to the log file
func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.logger != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
		}

		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		c.logger.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	}
}
