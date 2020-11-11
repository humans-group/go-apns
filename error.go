package apns

const (
	// Error when token is expired.
	ErrExpiredToken Error = "ExpiredProviderToken"
)

type Error string

func (e Error) Error() string {
	return string(e)
}
