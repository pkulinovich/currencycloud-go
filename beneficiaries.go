package currencycloud_go

import (
	"context"
	"fmt"
	"net/http"
)

// GetBeneficiary returns a beneficiary
func (c *Client) GetBeneficiary(ctx context.Context, id string) (*BeneficiaryResponse, error) {
	resp := &BeneficiaryResponse{}

	url := fmt.Sprintf(c.applyApiBaseUrl(EndpointGetBeneficiary), id)

	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return resp, err
	}

	err = c.SendWithAuth(req, resp)

	return resp, err
}

// CreateBeneficiary creates a new beneficiary.
func (c *Client) CreateBeneficiary(ctx context.Context, data *CreateBeneficiaryRequest) (*BeneficiaryResponse, error) {
	resp := &BeneficiaryResponse{}

	req, err := c.NewRequest(ctx, http.MethodPost, c.applyApiBaseUrl(EndpointCreateBeneficiary), data)
	if err != nil {
		return resp, err
	}

	err = c.SendWithAuth(req, resp)

	return resp, err
}

// UpdateBeneficiary updates a beneficiary
func (c *Client) UpdateBeneficiary(ctx context.Context, id string, data *UpdateBeneficiaryRequest) (*BeneficiaryResponse, error) {
	resp := &BeneficiaryResponse{}

	url := fmt.Sprintf(c.applyApiBaseUrl(EndpointUpdateBeneficiary), id)

	req, err := c.NewRequest(ctx, http.MethodPost, url, data)
	if err != nil {
		return resp, err
	}

	err = c.SendWithAuth(req, resp)

	return resp, err
}

// FindBeneficiaries find beneficiaries attached to the account or any sub-account owned by the authenticating user.
func (c *Client) FindBeneficiaries(ctx context.Context, params *FindBeneficiariesRequest) (*FindBeneficiariesResponse, error) {
	response := &FindBeneficiariesResponse{}

	req, err := c.NewRequest(ctx, http.MethodGet, c.applyApiBaseUrl(EndpointFindBeneficiaries), nil)
	if err != nil {
		return response, err
	}

	q, err := QueryStruct(params)
	if err != nil {
		return response, err
	}
	req.URL.RawQuery = q.Encode()

	err = c.SendWithAuth(req, response)

	return response, err
}

// DeleteBeneficiary remove a beneficiary
func (c *Client) DeleteBeneficiary(ctx context.Context, id string) (*BeneficiaryResponse, error) {
	resp := &BeneficiaryResponse{}

	url := fmt.Sprintf(c.applyApiBaseUrl(EndpointDeleteBeneficiary), id)

	req, err := c.NewRequest(ctx, http.MethodPost, url, nil)
	if err != nil {
		return resp, err
	}

	err = c.SendWithAuth(req, resp)

	return resp, err
}
