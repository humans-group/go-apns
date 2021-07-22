package apns

import (
	"github.com/sideshow/apns2"
)

const (
	DevelopmentGateway = apns2.HostDevelopment
	ProductionGateway  = apns2.HostProduction
)

// JWTConfig a set of data required for JWT authorization.
type JWTConfig struct {
	AuthKey []byte
	KeyID   string
	TeamID  string
}
