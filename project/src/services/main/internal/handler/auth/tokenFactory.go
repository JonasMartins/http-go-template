package auth

import (
	"time"
)

type TokenFactory interface {
	GenerateToken(email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*TokenPayload, error)
}
