package currencycloud_go_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	cc "github.com/pkulinovich/currencycloud-go"
	"net/http"
)

var _ = Describe("Authorization", func() {
	var (
		server *ghttp.Server
		client *cc.Client
		err    error

		loginID = "email@gmail.com"
		apiKey  = "19398ade0882ca5e0317ae5612bc3c939a21f9b5fd0f35442277d143f1c8b8d8"
		env     = "demonstration"

		token = "9307e7f58099731aa731ed3a6818807b"
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client, _ = cc.NewClient(loginID, apiKey, env)
		client.SetUrl(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Login", func() {
		var statusCode int
		var response interface{}
		var result *cc.AuthTokenResponse

		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, string(cc.EndpointLogin)),
					ghttp.VerifyHeader(http.Header{
						"Accept":       []string{cc.HeaderContentTypeApplicationJson},
						"User-Agent":   []string{cc.HeaderUserAgent},
						"Content-type": []string{cc.HeaderFormUrlencoded},
					}),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &response),
				),
			)
		})

		JustBeforeEach(func() {
			result, err = client.Login(context.Background())
		})

		Context("when successfully login", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
				response = cc.AuthTokenResponse{
					AuthToken: token,
				}
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should populate the token correctly", func() {
				Expect(result.AuthToken).To(Equal(token))
			})
		})

		Context("when the response fails to authenticate", func() {
			BeforeEach(func() {
				statusCode = http.StatusUnauthorized
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
			It("should return an empty AuthToken", func() {
				Expect(result.AuthToken).Should(BeEmpty())
			})
		})
	})

	Describe("Logout", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, string(cc.EndpointLogout)),
					ghttp.VerifyHeader(http.Header{
						"Accept":       []string{cc.HeaderContentTypeApplicationJson},
						"User-Agent":   []string{cc.HeaderUserAgent},
						"Content-type": []string{cc.HeaderContentTypeApplicationJson},
						"X-Auth-Token": []string{token},
					}),
				),
			)
		})

		JustBeforeEach(func() {
			client.SetAuthToken(token)
			err = client.Logout(context.Background())
		})

		Context("when the request succeeds", func() {
			It("should make a request to Logout without erroring", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
			})
		})
	})
})
