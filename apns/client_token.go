package apns

import (
	"fmt"

	"github.com/sideshow/apns2/token"
)

type Token struct {
	t *token.Token
}

func NewToken(keyPem []byte, keyID, teamID string) (*Token, error) {
	key, err := token.AuthKeyFromBytes(keyPem)
	if err != nil {
		return nil, fmt.Errorf("decode key: %w", err)
	}

	t := token.Token{
		AuthKey: key,
		KeyID:   keyID,
		TeamID:  teamID,
	}

	if _, err := t.Generate(); err != nil {
		return nil, fmt.Errorf("generate initial key: %w", err)
	}

	return &Token{
		t: &t,
	}, nil
}

func (t *Token) Bearer() (bearer string) {
	t.t.Lock()
	defer t.t.Unlock()

	if t.t.Expired() {
		_, _ = t.t.Generate()
	}

	return t.t.Bearer
}
