package apns

const (
	// API errors
	// https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/handling_notification_responses_from_apns
	errReasonExpiredProviderToken = "ExpiredProviderToken"

	// Error when token is expired.
	ErrExpiredToken Error = errReasonExpiredProviderToken
)

type Error string

func (e Error) Error() string {
	return string(e)
}
