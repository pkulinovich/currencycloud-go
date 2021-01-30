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

	CreateBeneficiaryRequest struct {
		Name                           string   `json:"name,omitempty"`
		BankAccountHolderName          string   `json:"bank_account_holder_name,omitempty"`
		BankCountry                    string   `json:"bank_country,omitempty"`
		Currency                       string   `json:"currency,omitempty"`
		Email                          string   `json:"email,omitempty"`
		BeneficiaryAddress             string   `json:"beneficiary_address,omitempty"`
		BeneficiaryCountry             string   `json:"beneficiary_country,omitempty"`
		AccountNumber                  string   `json:"account_number,omitempty"`
		RoutingCodeType1               string   `json:"routing_code_type_1,omitempty"`
		RoutingCodeValue1              string   `json:"routing_code_value_1,omitempty"`
		RoutingCodeType2               string   `json:"routing_code_type_2,omitempty"`
		RoutingCodeValue2              string   `json:"routing_code_value_2,omitempty"`
		BicSwift                       string   `json:"bic_swift,omitempty"`
		Iban                           string   `json:"iban,omitempty"`
		DefaultBeneficiary             string   `json:"default_beneficiary,omitempty"`
		BankAddress                    string   `json:"bank_address,omitempty"`
		BankName                       string   `json:"bank_name,omitempty"`
		BankAccountType                string   `json:"bank_account_type,omitempty"`
		BeneficiaryEntityType          string   `json:"beneficiary_entity_type,omitempty"`
		BeneficiaryCompanyName         string   `json:"beneficiary_company_name,omitempty"`
		BeneficiaryFirstName           string   `json:"beneficiary_first_name,omitempty"`
		BeneficiaryLastName            string   `json:"beneficiary_last_name,omitempty"`
		BeneficiaryCity                string   `json:"beneficiary_city,omitempty"`
		BeneficiaryPostcode            string   `json:"beneficiary_postcode,omitempty"`
		BeneficiaryStateOrProvince     string   `json:"beneficiary_state_or_province,omitempty"`
		BeneficiaryDateOfBirth         string   `json:"beneficiary_date_of_birth,omitempty"`
		BeneficiaryIdentificationType  string   `json:"beneficiary_identification_type,omitempty"`
		BeneficiaryIdentificationValue string   `json:"beneficiary_identification_value,omitempty"`
		PaymentTypes                   []string `json:"payment_types,omitempty"`
		OnBehalfOf                     string   `json:"on_behalf_of,omitempty"`
		BeneficiaryExternalReference   string   `json:"beneficiary_external_reference,omitempty"`
	}

	BeneficiaryResponse struct {
		ID                             string   `json:"id,omitempty"`
		BankAccountHolderName          string   `json:"bank_account_holder_name,omitempty"`
		Name                           string   `json:"name,omitempty"`
		Email                          string   `json:"email,omitempty"`
		PaymentTypes                   []string `json:"payment_types,omitempty"`
		BeneficiaryAddress             []string `json:"beneficiary_address,omitempty"`
		BeneficiaryCountry             string   `json:"beneficiary_country,omitempty"`
		BeneficiaryEntityType          string   `json:"beneficiary_entity_type,omitempty"`
		BeneficiaryCompanyName         string   `json:"beneficiary_company_name,omitempty"`
		BeneficiaryFirstName           string   `json:"beneficiary_first_name,omitempty"`
		BeneficiaryLastName            string   `json:"beneficiary_last_name,omitempty"`
		BeneficiaryCity                string   `json:"beneficiary_city,omitempty"`
		BeneficiaryPostcode            string   `json:"beneficiary_postcode,omitempty"`
		BeneficiaryStateOrProvince     string   `json:"beneficiary_state_or_province,omitempty"`
		BeneficiaryDateOfBirth         string   `json:"beneficiary_date_of_birth,omitempty"`
		BeneficiaryIdentificationType  string   `json:"beneficiary_identification_type,omitempty"`
		BeneficiaryIdentificationValue string   `json:"beneficiary_identification_value,omitempty"`
		BankCountry                    string   `json:"bank_country,omitempty"`
		BankName                       string   `json:"bank_name,omitempty"`
		BankAccountType                string   `json:"bank_account_type,omitempty"`
		Currency                       string   `json:"currency,omitempty"`
		AccountNumber                  string   `json:"account_number,omitempty"`
		RoutingCodeType1               string   `json:"routing_code_type_1,omitempty"`
		RoutingCodeValue1              string   `json:"routing_code_value_1,omitempty"`
		RoutingCodeType2               string   `json:"routing_code_type_2,omitempty"`
		RoutingCodeValue2              string   `json:"routing_code_value_2,omitempty"`
		BicSwift                       string   `json:"bic_swift,omitempty"`
		Iban                           string   `json:"iban,omitempty"`
		DefaultBeneficiary             string   `json:"default_beneficiary,omitempty"`
		CreatorContactId               string   `json:"creator_contact_id,omitempty"`
		BankAddress                    []string `json:"bank_address,omitempty"`
		CreatedAt                      string   `json:"created_at,omitempty"`
		UpdatedAt                      string   `json:"updated_at,omitempty"`
		BeneficiaryExternalReference   string   `json:"beneficiary_external_reference,omitempty"`
	}

	FindBeneficiariesRequest struct {
		OnBehalfOf                 string `json:"on_behalf_of,omitempty"`
		BankAccountHolderName      string `json:"bank_account_holder_name,omitempty"`
		BeneficiaryCountry         string `json:"beneficiary_country,omitempty"`
		Currency                   string `json:"currency,omitempty"`
		AccountNumber              string `json:"account_number,omitempty"`
		BicSwift                   string `json:"bic_swift,omitempty"`
		Iban                       string `json:"iban,omitempty"`
		DefaultBeneficiary         string `json:"default_beneficiary,omitempty"`
		BankName                   string `json:"bank_name,omitempty"`
		BankNameAccountType        string `json:"bank_account_type,omitempty"`
		Name                       string `json:"name,omitempty"`
		BeneficiaryEntityType      string `json:"beneficiary_entity_type,omitempty"`
		BeneficiaryCompanyName     string `json:"beneficiary_company_name,omitempty"`
		BeneficiaryFirstName       string `json:"beneficiary_first_name,omitempty"`
		BeneficiaryLastName        string `json:"beneficiary_last_name,omitempty"`
		BeneficiaryCity            string `json:"beneficiary_city,omitempty"`
		BeneficiaryPostcode        string `json:"beneficiary_postcode,omitempty"`
		BeneficiaryStateOrProvince string `json:"beneficiary_state_or_province,omitempty"`
		BeneficiaryDateOfBirth     string `json:"beneficiary_date_of_birth,omitempty"`
		Scope                      string `json:"scope,omitempty"`
		Page                       uint64 `json:"page,omitempty"`
		PerPage                    uint64 `json:"per_page,omitempty"`
		Order                      string `json:"order,omitempty"`
		OrderAscDesc               string `json:"order_asc_desc,omitempty"`
	}

	FindBeneficiariesResponse struct {
		Beneficiaries []*BeneficiaryResponse `json:"beneficiaries,omitempty"`
		Pagination    *Pagination            `json:"pagination,omitempty"`
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
