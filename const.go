package currencycloud_go

import "time"

type Endpoint string

const (
	EndpointLogin             Endpoint = "/authenticate/api"
	EndpointLogout            Endpoint = "/authenticate/close_session"
	EndpointGetBalance        Endpoint = "/balances/%s"
	EndpointFindBalances      Endpoint = "/balances/find"
	EndpointTopUpBalance      Endpoint = "/balances/top_up_margin"
	EndpointGetBeneficiary    Endpoint = "/beneficiaries/%s"
	EndpointDeleteBeneficiary Endpoint = "/beneficiaries/%s/delete"
	EndpointUpdateBeneficiary Endpoint = "/beneficiaries/%s"
	EndpointCreateBeneficiary Endpoint = "/beneficiaries/create"
	EndpointFindBeneficiaries Endpoint = "/beneficiaries/find"
	EndpointCreatePayment     Endpoint = "/payments/create"
	EndpointDeletePayment     Endpoint = "/payments/%s/delete"
	EndpointGetPayment        Endpoint = "/payments/%s"
	EndpointUpdatePayment     Endpoint = "/payments/%s"
)

type Environment string

const (
	EnvironmentProduction    Environment = "prod"
	EnvironmentDemonstration Environment = "demonstration"
	EnvironmentUAT           Environment = "uat"
)

const (
	HeaderContentTypeApplicationJson string = "application/json"
	HeaderFormUrlencoded             string = "application/x-www-form-urlencoded"
	HeaderUserAgent                  string = "CurrencyCloud-Go-http-client/2.0"
)

const (
	ExpiresInLimit = time.Duration(60) * time.Second
)
