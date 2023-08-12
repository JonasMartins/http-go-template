package auth

import (
	"errors"
	"project/src/pkg/utils"
	"time"
)

type TokenPayload struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewTokenPayload(email string, duration time.Duration) (*TokenPayload, error) {
	newUUID, err := utils.GenerateNewUUid()
	if err != nil {
		return nil, err
	}
	res := &TokenPayload{
		ID:        string(newUUID),
		Email:     email,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return res, nil
}

func (tr *TokenPayload) Valid() error {
	if time.Now().After(tr.ExpiresAt) {
		return errors.New("token expired")
	}
	return nil
}
