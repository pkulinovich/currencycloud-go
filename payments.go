package currencycloud_go

import (
	"context"
	"fmt"
	"net/http"
)

// GetPayment retrieve a payment record.
func (c *Client) GetPayment(ctx context.Context, id string) (*PaymentResponse, error) {
	resp := &PaymentResponse{}

	url := fmt.Sprintf(c.applyApiBaseUrl(EndpointGetPayment), id)

	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return resp, err
	}

	err = c.SendWithAuth(req, resp)

	return resp, err
}

// CreatePayment make a payment.
func (c *Client) CreatePayment(ctx context.Context, data *CreatePaymentRequest) (*PaymentResponse, error) {
	resp := &PaymentResponse{}

	req, err := c.NewRequest(ctx, http.MethodPost, c.applyApiBaseUrl(EndpointCreatePayment), data)
	if err != nil {
		return resp, err
	}

	err = c.SendWithAuth(req, resp)

	return resp, err
}
