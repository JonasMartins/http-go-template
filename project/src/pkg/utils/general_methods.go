package utils

import (
	"bytes"
	"log"
	"os/exec"

	"golang.org/x/crypto/bcrypt"
)

// * Generates a new valid uuid as []byte
func GenerateNewUUid() ([]byte, error) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(newUUID, []byte("\n")), nil
}

// * Returns a fatalf response
func FatalResult(s string, err error) {
	log.Fatalf("%s %v", s, err)
}

// * return nil if pass and hash are the same, error otherwise
func ValidatePassword(pass string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
