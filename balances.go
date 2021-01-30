package currencycloud_go

import (
	"context"
	"fmt"
	"net/http"
)

// GetBalance gets the balance for a currency from the account of the authenticating user.
func (c *Client) GetBalance(ctx context.Context, currency string) (*BalanceResponse, error) {
	resp := &BalanceResponse{}

	url := fmt.Sprintf(c.applyApiBaseUrl(EndpointGetBalance), currency)

	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return resp, err
	}

	if err = c.SendWithAuth(req, resp); err != nil {
		return resp, err
	}

	return resp, nil
}

// Find search for currency balances in the main account or a sub-account owned by the authenticating user.
func (c *Client) FindBalances(ctx context.Context) (*BalancesListResponse, error) {
	resp := &BalancesListResponse{}

	req, err := c.NewRequest(ctx, http.MethodGet, c.applyApiBaseUrl(EndpointFindBalances), nil)
	if err != nil {
		return resp, err
	}

	if err = c.SendWithAuth(req, resp); err != nil {
		return resp, err
	}

	return resp, nil
}

// TopUpMargin tops up the Margin Balance
func (c *Client) TopUpMarginBalance(ctx context.Context, data *BalanceTopUpMarginRequest) (*BalanceTopUpMarginResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, c.applyApiBaseUrl(EndpointTopUpBalance), data)
	response := &BalanceTopUpMarginResponse{}
	if err != nil {
		return response, err
	}
	err = c.SendWithAuth(req, response)
	return response, err
}
