package currencycloud_go

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type (
	Client struct {
		sync.Mutex
		client         *http.Client
		loginID        string
		apiKey         string
		env            Environment
		logger         io.Writer
		authToken      *AuthTokenResponse
		tokenExpiresAt time.Time
	}

	AuthTokenResponse struct {
		AuthToken string `json:"auth_token"`
	}

	BalanceRequest struct {
		Currency string `json:"currency"`
	}

	BalanceTopUpMarginRequest struct {
		Currency   string `json:"currency"`
		Amount     string `json:"amount"`
		OnBehalfOf string `json:"on_behalf_of,omitempty"`
	}

	BalanceResponse struct {
		ID        string `json:"id,omitempty"`
		AccountID string `json:"account_id,omitempty"`
		Currency  string `json:"currency,omitempty"`
		Amount    string `json:"amount,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	}

	BalancesListResponse struct {
		Balances   []*BalanceResponse `json:"balances,omitempty"`
		Pagination *Pagination        `json:"pagination,omitempty"`
	}

	BalanceTopUpMarginResponse struct {
		AccountID         string `json:"account_id,omitempty"`
		Currency          string `json:"currency,omitempty"`
		TransferredAmount string `json:"transferred_amount,omitempty"`
	}

	Pagination struct {
		TotalEntries uint64 `json:"total_entries,omitempty"`
		TotalPages   uint64 `json:"total_pages,omitempty"`
		CurrentPage  uint64 `json:"current_page,omitempty"`
		PerPage      uint64 `json:"per_page,omitempty"`
		PreviousPage int64  `json:"previous_page,omitempty"`
		NextPage     int64  `json:"next_page,omitempty"`
		Order        string `json:"order,omitempty"`
		OrderAscDesc string `json:"order_asc_desc,omitempty"`
	}

	ErrorMessage struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
		Params  string `json:"params,omitempty"`
	}

	ErrorResponse struct {
		Response      *http.Response             `json:"-"`
		ErrorCode     string                     `json:"error_code,omitempty"`
		ErrorMessages map[string][]*ErrorMessage `json:"error_messages,omitempty"`
	}
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s, %+v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.ErrorCode, r.ErrorMessages)
}
