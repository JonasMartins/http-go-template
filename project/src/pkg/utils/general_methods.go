package utils

import (
	"log"
	"os/exec"
)

// * Generates a new valid uuid as []byte
func GenerateNewUUid() ([]byte, error) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, err
	}
	return newUUID, nil
}

// * Returns a fatalf response
func FatalResult(s string, err error) {
	log.Fatalf("%s %v", s, err)
}
