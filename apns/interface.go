package apns

import "context"

// Client describes methods used to send apns notifications.
type Client interface {
	Send(ctx context.Context, n *Notification) error
}
