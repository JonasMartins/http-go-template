package auth

import (
	"errors"
	"project/src/pkg/utils"
	"time"
)

type TokenFactory interface {
	GenerateToken(email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*TokenResponse, error)
}

type TokenResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewTokenResponse(email string, duration time.Duration) (*TokenResponse, error) {
	newUUID, err := utils.GenerateNewUUid()
	if err != nil {
		return nil, err
	}
	res := &TokenResponse{
		ID:        string(newUUID),
		Email:     email,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return res, nil
}

func (tr *TokenResponse) Valid() error {
	if time.Now().After(tr.ExpiresAt) {
		return errors.New("token expired")
	}
	return nil
}
