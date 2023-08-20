package test

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {

	pass := "pb_admin"
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		t.Logf("Err %v", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		t.Logf("Err %v", err)
	}
}
