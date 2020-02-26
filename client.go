package apns

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/http2"
)

const (
	DevelopmentGateway = "https://api.development.push.apple.com"
	ProductionGateway  = "https://api.push.apple.com"
)

type Client interface {
	Send(ctx context.Context, n *Notification) error
}

var _ Client = &SimpleClient{}

// SimpleClient represents the Apple Push Notification Service that you send notifications to.
type SimpleClient struct {
	http     *http.Client
	endpoint string
}

// NewClient creates new APNS client based on defined Options.
func NewClient(opts ...ClientOption) *SimpleClient {
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
		return errors.Wrap(err, "failed to prepare http req")
	}

	if err := c.do(req); err != nil {
		return errors.Wrap(err, "failed to do http request")
	}

	return nil
}

func (c *SimpleClient) prepareRequest(ctx context.Context, n *Notification) (*http.Request, error) {
	data, err := n.Payload.MarshalJSON()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal notification payload %#v", n)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
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
	resp, err := c.http.Do(req)
	if err != nil {
		return errors.Wrapf(err, "failed to do http request %#v", *req)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to read full response body %s", string(respBytes))
	}

	rawResp := new(RawResp)
	if err := rawResp.UnmarshalJSON(respBytes); err != nil {
		return errors.Wrapf(err, "failed to unmarshal response %s", string(respBytes))
	}

	if rawResp.Reason != "" {
		errors.New(rawResp.Reason)
	}

	return nil
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
