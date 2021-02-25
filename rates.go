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
