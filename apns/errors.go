package apns

import (
	"fmt"
)

const (
	ErrExpiredToken   Error = "ExpiredProviderToken"
	ErrBadDeviceToken Error = "BadDeviceToken"
	ErrUnregistered   Error = "Unregistered"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

// Map API error reason to client error if there is a reason.
func apiErrorReasonToClientError(reason ErrorReason) error {
	switch reason {
	case "":
		return nil
	case ReasonExpiredProviderToken:
		return ErrExpiredToken
	case ReasonBadDeviceToken:
		return ErrBadDeviceToken
	case ReasonCodeUnregistered:
		return ErrUnregistered
	default:
		return fmt.Errorf(string(reason))
	}
}
