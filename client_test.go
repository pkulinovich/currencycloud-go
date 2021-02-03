package currencycloud_go_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cc "github.com/pkulinovich/currencycloud-go"
	"net/url"
)

var _ = Describe("Client", func() {
	var (
		loginID string
		apiKey  string
		env     string

		client *cc.Client
		err    error
	)

	BeforeEach(func() {
		loginID = "email@gmail.com"
		apiKey = "19398ade0882ca5e0317ae5612bc3c939a21f9b5fd0f35442277d143f1c8b8d8"
		env = "demonstration"

		client, err = cc.NewClient(loginID, apiKey, env)
	})

	Describe("init new client", func() {
		Context("when the credentials are correct", func() {
			It("should populate the arguments correctly", func() {
				Expect(client.LoginID).To(Equal(loginID))
				Expect(client.ApiKey).To(Equal(apiKey))
				Expect(client.Env).To(Equal(cc.EnvironmentDemonstration))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("extracting the credentials", func() {
		It("should correctly return the credentials", func() {
			v := url.Values{}
			v.Set("login_id", loginID)
			v.Set("api_key", apiKey)

			Expect(client.GetCredentials().Encode()).To(Equal(v.Encode()))
		})
	})
})
