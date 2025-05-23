package apns

import (
	. "github.com/onsi/ginkgo/v2"
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
				errorReason = ReasonExpiredProviderToken
				err := apiErrorReasonToClientError(errorReason)
				Ω(err).Should(Equal(ErrExpiredToken))
			})
		})

		When("error reason is reasonBadDeviceToken", func() {
			It("should return ErrBadDeviceToken", func() {
				errorReason = ReasonBadDeviceToken
				err := apiErrorReasonToClientError(errorReason)
				Ω(err).Should(Equal(ErrBadDeviceToken))
			})
		})

		When("error reason is reasonCodeUnregistered", func() {
			It("should return ErrUnregistered", func() {
				errorReason = ReasonCodeUnregistered
				err := apiErrorReasonToClientError(errorReason)
				Ω(err).Should(Equal(ErrUnregistered))
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
