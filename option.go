package apns

import (
	"crypto/tls"
	"errors"
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

// WithEndpoint specifies custom APN endpoint. Useful for test propose.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *SimpleClient) error {
		c.endpoint = endpoint
		return nil
	}
}

// WithCertificate is Option to configure TLS certificates for HTTP connection.
// Certificates should be used with BundleID, that is possible to set by
// `WithBundleID` option.
func WithCertificate(crt tls.Certificate) ClientOption {
	return func(c *SimpleClient) error {
		config := &tls.Config{
			Certificates: []tls.Certificate{crt},
		}
		config.BuildNameToCertificate()
		c.http.Transport.(*http2.Transport).TLSClientConfig = config
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

// WithMaxIdleConnections sets maximum number of the idle HTTP connection
// that can be reused in order do not create new connection.
func WithMaxIdleConnections(maxIdleConn int) ClientOption {
	return func(c *SimpleClient) error {
		if maxIdleConn < 1 {
			return errors.New("invalid MaxIdleConnsPerHost")
		}
		c.http.Transport.(*http.Transport).MaxIdleConnsPerHost = maxIdleConn
		return nil
	}
}
