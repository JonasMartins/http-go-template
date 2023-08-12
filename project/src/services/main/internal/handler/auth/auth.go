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
	payload, err := NewTokenPayload(email, duration)
	if err != nil {
		return "", err
	}
	return f.paseto.Encrypt(f.key, payload, nil)
}
func (f *PasetoFactory) VerifyToken(token string) (*TokenPayload, error) {
	payload := &TokenPayload{}

	err := f.paseto.Decrypt(token, f.key, payload, nil)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
