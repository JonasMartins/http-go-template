package utils

import "os/exec"

func GenerateNewUUid() ([]byte, error) {

	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, err
	}
	return newUUID, nil
}
