package apns

import (
	"fmt"
	"net/http"

	"github.com/sideshow/apns2/token"
)

// RoundTripperJWTDecorator an implementation of http.RoundTripper interface
// with ability to specify authorization JWT token for each request.
type RoundTripperJWTDecorator struct {
	Parent http.RoundTripper
	Token  *token.Token
}

func (d *RoundTripperJWTDecorator) RoundTrip(r *http.Request) (*http.Response, error) {
	d.Token.GenerateIfExpired()
	r.Header.Set("Authorization", fmt.Sprintf("bearer %s", d.Token.Bearer))

	return d.Parent.RoundTrip(r)
}
