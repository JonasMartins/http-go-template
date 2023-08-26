package test

import (
	"log"
	"project/src/services/main/configs"
	"testing"
)

func TestGinLogger(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	f, err := configs.GinLogger()
	if err != nil {
		t.Logf("Err %v", err)
	}
	t.Log(f)
}
