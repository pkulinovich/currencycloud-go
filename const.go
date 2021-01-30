package currencycloud_go

import "time"

type Endpoint string

const (
	EndpointLogin             Endpoint = "authenticate/api"
	EndpointGetBalance        Endpoint = "balances/%s"
	EndpointFindBalances      Endpoint = "balances/find"
	EndpointTopUpBalance      Endpoint = "balances/top_up_margin"
	EndpointGetBeneficiary    Endpoint = "beneficiaries/%s"
	EndpointCreateBeneficiary Endpoint = "beneficiaries/create"
	EndpointFindBeneficiaries Endpoint = "beneficiaries/find"
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