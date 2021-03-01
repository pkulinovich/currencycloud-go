package currencycloud_go

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetBasicRates(ctx context.Context, currencyPair string) (*RatesResponse, error) {
	resp := &RatesResponse{}

	url := fmt.Sprintf(c.applyApiBaseUrl(EndpointGetBasicRates), currencyPair)

	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return resp, err
	}

	err = c.SendWithAuth(req, resp)

	return resp, err
}

func (c *Client) GetDetailedRates(ctx context.Context, params *GetDetailedRatesRequest) (*DetailedRatesResponse, error) {
	resp := &DetailedRatesResponse{}

	url := fmt.Sprintf(c.applyApiBaseUrl(EndpointGetDetailedRates))

	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return resp, err
	}

	q, err := QueryStruct(params)
	if err != nil {
		return resp, err
	}
	req.URL.RawQuery = q.Encode()

	err = c.SendWithAuth(req, resp)

	return resp, err
}
