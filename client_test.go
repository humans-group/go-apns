package apns

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	Context("apiErrorReasonToClientError func", func() {
		var errorReason ErrorReason

		When("error reason is empty", func() {
			It("should succeed", func() {
				errorReason = ""
				err := apiErrorReasonToClientError(errorReason)
				Ω(err).Should(Succeed())
			})
		})

		When("error reason is ExpiredProviderToken", func() {
			It("should return ErrExpiredToken", func() {
				errorReason = reasonExpiredProviderToken
				err := apiErrorReasonToClientError(errorReason)
				Ω(err).Should(Equal(ErrExpiredToken))
			})
		})

		When("error reason is general error", func() {
			It("should return Error", func() {
				errorReason = "anyOtherError"
				err := apiErrorReasonToClientError(errorReason)
				Ω(err).Should(HaveOccurred())
				Ω(err.Error()).Should(Equal(string(errorReason)))
			})
		})
	})
})
