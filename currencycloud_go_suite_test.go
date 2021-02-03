package currencycloud_go_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCurrencycloudGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CurrencycloudGo Suite")
}
