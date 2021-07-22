package apns

import (
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

// ClientOption defines athe APNS SimpleClient option.
type ClientOption func(c *SimpleClient) error

// WithHTTPClient sets custom HTTP SimpleClient.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *SimpleClient) error {
		c.http = httpClient
		return nil
	}
}

// WithGateway specifies custom APN endpoint. Useful for test propose.
func WithGateway(endpoint string) ClientOption {
	return func(c *SimpleClient) error {
		c.endpoint = endpoint
		return nil
	}
}

// WithTimeout sets HTTP SimpleClient timeout.
func WithTimeout(t time.Duration) ClientOption {
	return func(c *SimpleClient) error {
		c.http.Timeout = t
		return nil
	}
}

// WithJWTAuthorization enables JWT authorization for apns requests.
func WithJWTAuthorization(cfg JWTConfig) ClientOption {
	return func(c *SimpleClient) error {
		authToken, err := NewToken(cfg.AuthKey, cfg.KeyID, cfg.TeamID)
		if err != nil {
			return err
		}

		parent := c.http.Transport
		if parent == nil {
			parent = &http2.Transport{}
		}

		c.http.Transport = &RoundTripperJWTDecorator{
			Parent: parent,
			Token:  authToken,
		}
		return nil
	}
}
