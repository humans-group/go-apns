package apns

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sideshow/apns2/token"
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
		authKey, err := token.AuthKeyFromBytes(cfg.AuthKey)
		if err != nil {
			return fmt.Errorf("parse auth key: %w", err)
		}

		parent := c.http.Transport
		if parent == nil {
			parent = &http2.Transport{}
		}

		c.http.Transport = &RoundTripperJWTDecorator{
			Parent: parent,
			Token: &token.Token{
				AuthKey: authKey,
				KeyID:   cfg.KeyID,
				TeamID:  cfg.TeamID,
			},
		}
		return nil
	}
}
