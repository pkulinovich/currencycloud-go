package currencycloud_go

var urls = map[Environment]string{
	EnvironmentProduction:    "https://api.currencycloud.com/v2/",
	EnvironmentDemonstration: "https://devapi.currencycloud.com/v2/",
	EnvironmentUAT:           "https://api-uat1.ccycloud.com/",
}
