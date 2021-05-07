package apns

const (
	// Error when token is expired.
	ErrExpiredToken   Error = "ExpiredProviderToken"
	ErrBadDeviceToken Error = "BadDeviceToken"
	ErrUnregistered   Error = "Unregistered"
)

type Error string

func (e Error) Error() string {
	return string(e)
}
