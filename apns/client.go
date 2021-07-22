package apns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

var _ Client = (*SimpleClient)(nil)

// SimpleClient represents the Apple Push Notification Service that you send notifications to.
type SimpleClient struct {
	http     *http.Client
	endpoint string
}

// MustNewClient creates new APNS client based on defined Options.
func MustNewClient(opts ...ClientOption) *SimpleClient {
	c := &SimpleClient{
		http: &http.Client{
			Transport: &http2.Transport{},
		},
		endpoint: ProductionGateway,
	}
	for _, o := range opts {
		if err := o(c); err != nil {
			panic(fmt.Sprintf("failed to apply opt: %v", err))
		}
	}

	c.endpoint = fmt.Sprintf("%s/3/device/", c.endpoint)

	return c
}

// Send sends Notification to the APN service.
func (c *SimpleClient) Send(ctx context.Context, n *Notification) error {
	req, err := c.prepareRequest(ctx, n)
	if err != nil {
		return fmt.Errorf("prepare request: %w", err)
	}

	return c.do(req)
}

func (c *SimpleClient) prepareRequest(ctx context.Context, n *Notification) (*http.Request, error) {
	data, err := n.Payload.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("marshal payload json: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.endpoint+n.DeviceToken,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}

	setHeaders(req, n)

	return req, nil
}

func (c *SimpleClient) do(req *http.Request) error {
	httpResp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer func() { _ = httpResp.Body.Close() }()

	if httpResp.StatusCode == http.StatusOK {
		return nil
	}

	var resp Response
	err = json.
		NewDecoder(httpResp.Body).
		Decode(&resp)
	if err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return apiErrorReasonToClientError(resp.Reason)
}

func setHeaders(r *http.Request, n *Notification) {
	r.Header.Set("Content-Type", "application/json; charset=utf-8")

	if n.Topic != "" {
		r.Header.Set("apns-topic", n.Topic)
	}

	if n.ApnsID != "" {
		r.Header.Set("apns-id", n.ApnsID)
	}

	if n.CollapseID != "" {
		r.Header.Set("apns-collapse-id", n.CollapseID)
	}

	if n.Priority > 0 {
		r.Header.Set("apns-priority", fmt.Sprintf("%v", n.Priority))
	}

	if !n.Expiration.IsZero() {
		r.Header.Set("apns-expiration", fmt.Sprintf("%v", n.Expiration.Unix()))
	}

	if n.PushType != "" {
		r.Header.Set("apns-push-type", string(n.PushType))
	} else {
		r.Header.Set("apns-push-type", string(PushTypeAlert))
	}
}
