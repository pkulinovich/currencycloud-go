package currencycloud_go

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetBeneficiaryRequirements(
	ctx context.Context,
	params *GetBeneficiaryRequirementsRequest,
) (*GetBeneficiaryRequirementsResponse, error) {
	resp := &GetBeneficiaryRequirementsResponse{}

	url := fmt.Sprintf(c.applyApiBaseUrl(EndpointGetBeneficiaryRequirements))

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
