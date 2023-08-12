package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoFactory struct {
	paseto *paseto.V2
	key    []byte
}

func NewPasetoFactory(key string) (TokenFactory, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key length, must have %d characters", chacha20poly1305.KeySize)
	}

	obj := &PasetoFactory{
		paseto: paseto.NewV2(),
		key:    []byte(key),
	}

	return obj, nil
}

func (f *PasetoFactory) GenerateToken(email string, duration time.Duration) (string, error) {
	tr, err := NewTokenResponse(email, duration)
	if err != nil {
		return "", err
	}
	return f.paseto.Encrypt(f.key, tr, nil)
}
func (f *PasetoFactory) VerifyToken(token string) (*TokenResponse, error) {
	tr := &TokenResponse{}

	err := f.paseto.Decrypt(token, f.key, tr, nil)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	err = tr.Valid()
	if err != nil {
		return nil, err
	}
	return tr, nil
}
