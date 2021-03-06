package currencycloud_go

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
)

type (
	Client struct {
		sync.Mutex
		client         *http.Client
		LoginID        string
		ApiKey         string
		Env            Environment
		URL            string
		logger         io.Writer
		authToken      string
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
		Name                           string `json:"name,omitempty"`
		BankAccountHolderName          string `json:"bank_account_holder_name,omitempty"`
		BankCountry                    string `json:"bank_country,omitempty"`
		Currency                       string `json:"currency,omitempty"`
		Email                          string `json:"email,omitempty"`
		BeneficiaryAddress             string `json:"beneficiary_address,omitempty"`
		BeneficiaryCountry             string `json:"beneficiary_country,omitempty"`
		AccountNumber                  string `json:"account_number,omitempty"`
		RoutingCodeType1               string `json:"routing_code_type_1,omitempty"`
		RoutingCodeValue1              string `json:"routing_code_value_1,omitempty"`
		RoutingCodeType2               string `json:"routing_code_type_2,omitempty"`
		RoutingCodeValue2              string `json:"routing_code_value_2,omitempty"`
		BicSwift                       string `json:"bic_swift,omitempty"`
		Iban                           string `json:"iban,omitempty"`
		DefaultBeneficiary             string `json:"default_beneficiary,omitempty"`
		BankAddress                    string `json:"bank_address,omitempty"`
		BankName                       string `json:"bank_name,omitempty"`
		BankAccountType                string `json:"bank_account_type,omitempty"`
		BeneficiaryEntityType          string `json:"beneficiary_entity_type,omitempty"`
		BeneficiaryCompanyName         string `json:"beneficiary_company_name,omitempty"`
		BeneficiaryFirstName           string `json:"beneficiary_first_name,omitempty"`
		BeneficiaryLastName            string `json:"beneficiary_last_name,omitempty"`
		BeneficiaryCity                string `json:"beneficiary_city,omitempty"`
		BeneficiaryPostcode            string `json:"beneficiary_postcode,omitempty"`
		BeneficiaryStateOrProvince     string `json:"beneficiary_state_or_province,omitempty"`
		BeneficiaryDateOfBirth         string `json:"beneficiary_date_of_birth,omitempty"`
		BeneficiaryIdentificationType  string `json:"beneficiary_identification_type,omitempty"`
		BeneficiaryIdentificationValue string `json:"beneficiary_identification_value,omitempty"`
		PaymentTypes                   string `json:"payment_types[],omitempty"`
		OnBehalfOf                     string `json:"on_behalf_of,omitempty"`
		BeneficiaryExternalReference   string `json:"beneficiary_external_reference,omitempty"`
	}

	UpdateBeneficiaryRequest struct {
		Name                           string `json:"name,omitempty"`
		BankAccountHolderName          string `json:"bank_account_holder_name,omitempty"`
		BankCountry                    string `json:"bank_country,omitempty"`
		Currency                       string `json:"currency,omitempty"`
		Email                          string `json:"email,omitempty"`
		BeneficiaryAddress             string `json:"beneficiary_address,omitempty"`
		BeneficiaryCountry             string `json:"beneficiary_country,omitempty"`
		AccountNumber                  string `json:"account_number,omitempty"`
		RoutingCodeType1               string `json:"routing_code_type_1,omitempty"`
		RoutingCodeValue1              string `json:"routing_code_value_1,omitempty"`
		RoutingCodeType2               string `json:"routing_code_type_2,omitempty"`
		RoutingCodeValue2              string `json:"routing_code_value_2,omitempty"`
		BicSwift                       string `json:"bic_swift,omitempty"`
		Iban                           string `json:"iban,omitempty"`
		DefaultBeneficiary             string `json:"default_beneficiary,omitempty"`
		BankAddress                    string `json:"bank_address,omitempty"`
		BankName                       string `json:"bank_name,omitempty"`
		BankAccountType                string `json:"bank_account_type,omitempty"`
		BeneficiaryEntityType          string `json:"beneficiary_entity_type,omitempty"`
		BeneficiaryCompanyName         string `json:"beneficiary_company_name,omitempty"`
		BeneficiaryFirstName           string `json:"beneficiary_first_name,omitempty"`
		BeneficiaryLastName            string `json:"beneficiary_last_name,omitempty"`
		BeneficiaryCity                string `json:"beneficiary_city,omitempty"`
		BeneficiaryPostcode            string `json:"beneficiary_postcode,omitempty"`
		BeneficiaryStateOrProvince     string `json:"beneficiary_state_or_province,omitempty"`
		BeneficiaryDateOfBirth         string `json:"beneficiary_date_of_birth,omitempty"`
		BeneficiaryIdentificationType  string `json:"beneficiary_identification_type,omitempty"`
		BeneficiaryIdentificationValue string `json:"beneficiary_identification_value,omitempty"`
		PaymentTypes                   string `json:"payment_types[],omitempty"`
		OnBehalfOf                     string `json:"on_behalf_of,omitempty"`
		BeneficiaryExternalReference   string `json:"beneficiary_external_reference,omitempty"`
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

	BeneficiaryRequirements struct {
		PaymentType                string `json:"payment_type,omitempty"`
		BeneficiaryEntityType      string `json:"beneficiary_entity_type,omitempty"`
		BeneficiaryAddress         string `json:"beneficiary_address,omitempty"`
		BeneficiaryCity            string `json:"beneficiary_city,omitempty"`
		BeneficiaryCountry         string `json:"beneficiary_country,omitempty"`
		BeneficiaryFirstName       string `json:"beneficiary_first_name,omitempty"`
		BeneficiaryLastName        string `json:"beneficiary_last_name,omitempty"`
		BeneficiaryCompanyName     string `json:"beneficiary_company_name,omitempty"`
		BeneficiaryPostcode        string `json:"beneficiary_postcode,omitempty"`
		BeneficiaryStateOrProvince string `json:"beneficiary_state_or_province,omitempty"`
		BicSwift                   string `json:"bic_swift,omitempty"`
		Iban                       string `json:"iban,omitempty"`
		AcctNumber                 string `json:"acct_number,omitempty"`
		SortCode                   string `json:"sort_code,omitempty"`
		Aba                        string `json:"aba,omitempty"`
		BsbCode                    string `json:"bsb_code,omitempty"`
		InstitutionNo              string `json:"institution_no,omitempty"`
		BankCode                   string `json:"bank_code,omitempty"`
		BranchCode                 string `json:"branch_code,omitempty"`
		Clabe                      string `json:"clabe,omitempty"`
		Cnaps                      string `json:"cnaps,omitempty"`
		Ifsc                       string `json:"ifsc,omitempty"`
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

	CreatePaymentRequest struct {
		Currency      string `json:"currency,omitempty"`
		BeneficiaryID string `json:"beneficiary_id,omitempty"`
		Amount        string `json:"amount,omitempty"`
		Reason        string `json:"reason,omitempty"`
		Reference     string `json:"reference,omitempty"`
		PaymentType   string `json:"payment_type,omitempty"`
		PurposeCode   string `json:"purpose_code,omitempty"`
	}

	PaymentResponse struct {
		ID                         string `json:"id,omitempty"`
		Amount                     string `json:"amount,omitempty"`
		BeneficiaryID              string `json:"beneficiary_id,omitempty"`
		Currency                   string `json:"currency,omitempty"`
		Reference                  string `json:"reference,omitempty"`
		Reason                     string `json:"reason,omitempty"`
		Status                     string `json:"status,omitempty"`
		CreatorContactID           string `json:"creator_contact_id,omitempty"`
		PaymentType                string `json:"payment_type,omitempty"`
		PaymentDate                string `json:"payment_date,omitempty"`
		TransferredAt              string `json:"transferred_at,omitempty"`
		AuthorisationStepsRequired string `json:"authorisation_steps_required,omitempty"`
		LastUpdaterContactID       string `json:"last_updater_contact_id,omitempty"`
		ShortReference             string `json:"short_reference,omitempty"`
		ConversionID               string `json:"conversion_id,omitempty"`
		FailureReason              string `json:"failure_reason,omitempty"`
		PayerID                    string `json:"payer_id,omitempty"`
		PayerDetailsSource         string `json:"payer_details_source,omitempty"`
		CreatedAt                  string `json:"created_at,omitempty"`
		UpdatedAt                  string `json:"updated_at,omitempty"`
		PaymentGroupID             string `json:"payment_group_id,omitempty"`
		UniqueRequestID            string `json:"unique_request_id,omitempty"`
		FailureReturnedAmount      string `json:"failure_returned_amount,omitempty"`
		UltimateBeneficiaryName    string `json:"ultimate_beneficiary_name,omitempty"`
		PurposeCode                string `json:"purpose_code,omitempty"`
		ChargeType                 string `json:"charge_type,omitempty"`
		FeeAmount                  string `json:"fee_amount,omitempty"`
		FeeCurrency                string `json:"fee_currency,omitempty"`
	}

	GetDetailedRatesRequest struct {
		OnBehalfOf               string `json:"on_behalf_of,omitempty"`
		BuyCurrency              string `json:"buy_currency,omitempty"`
		SellCurrency             string `json:"sell_currency,omitempty"`
		FixedSide                string `json:"fixed_side,omitempty"`
		Amount                   string `json:"amount,omitempty"`
		ConversionDate           string `json:"conversion_date,omitempty"`
		ConversionDatePreference string `json:"conversion_date_preference,omitempty"`
	}

	DetailedRatesResponse struct {
		SettlementCutOffTime string `json:"settlement_cut_off_time,omitempty"`
		CurrencyPair         string `json:"currency_pair,omitempty"`
		ClientBuyCurrency    string `json:"client_buy_currency,omitempty"`
		ClientSellCurrency   string `json:"client_sell_currency,omitempty"`
		ClientBuyAmount      string `json:"client_buy_amount,omitempty"`
		ClientSellAmount     string `json:"client_sell_amount,omitempty"`
		FixedSide            string `json:"fixed_side,omitempty"`
		ClientRate           string `json:"client_rate,omitempty"`
		PartnerRate          string `json:"partner_rate,omitempty"`
		CoreRate             string `json:"core_rate,omitempty"`
		DepositRequired      bool   `json:"deposit_required,omitempty"`
		DepositAmount        string `json:"deposit_amount,omitempty"`
		DepositCurrency      string `json:"deposit_currency,omitempty"`
		MidMarketRate        string `json:"mid_market_rate,omitempty"`
	}

	RatesResponse struct {
		Rates       map[string][]string `json:"rates,omitempty"`
		Unavailable []string            `json:"unavailable,omitempty"`
	}

	GetBeneficiaryRequirementsRequest struct {
		Currency           string `json:"currency"`
		BankAccountCountry string `json:"bank_account_country"`
		BeneficiaryCountry string `json:"beneficiary_country,omitempty"`
	}

	GetBeneficiaryRequirementsResponse struct {
		Details []*BeneficiaryRequirements `json:"details"`
	}

	GetPaymentPurposeCodesRequest struct {
		Currency           string `json:"currency"`
		BankAccountCountry string `json:"bank_account_country"`
		EntityType         string `json:"entity_type,omitempty"`
	}

	GetPaymentPurposeCodesResponse struct {
		PurposeCodes []*PaymentPurposeCode `json:"purpose_codes"`
	}

	PaymentPurposeCode struct {
		BankAccountCountry string `json:"bank_account_country"`
		Currency           string `json:"currency"`
		EntityType         string `json:"entity_type"`
		PurposeCode        string `json:"purpose_code"`
		PurposeDescription string `json:"purpose_description"`
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
	res, _ := json.Marshal(r)
	return string(res)
}
