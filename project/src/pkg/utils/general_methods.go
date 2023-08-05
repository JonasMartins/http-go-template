package utils

import (
	"log"
	"os/exec"
)

func GenerateNewUUid() ([]byte, error) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, err
	}
	return newUUID, nil
}

func FatalResult(s string, err error) {
	log.Fatalf("%s %v", s, err)
}
