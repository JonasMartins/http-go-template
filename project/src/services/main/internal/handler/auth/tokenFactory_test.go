package auth

import (
	"log"
	"os"
	"project/src/pkg/utils"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// setup statements
	setup()

	// run the tests
	e := m.Run()

	// cleanup statements
	teardown()

	// report the exit code
	os.Exit(e)
}

func setup() {
	log.Println("Setting up.")
}

func teardown() {
	log.Println("Tearing down.")
}

func init() {
	log.Println("Init setup.")
}

func TestPasetoTokenGenerator(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	r := utils.RandomString(32)
	p, err := NewPasetoFactory(r)
	if err != nil {
		t.Errorf("error at generating object: %s", err.Error())
		return
	}
	email := "admin@email.com"
	duration := time.Minute

	token, err := p.GenerateToken(email, duration)
	if err != nil {
		t.Errorf("error at generating token: %s", err.Error())
		return
	}
	if len(token) == 0 {
		t.Errorf("token length invalid: %s", err.Error())
		return
	}

	payload, err := p.VerifyToken(token)
	if err != nil {
		t.Errorf("error at validating token: %s", err.Error())
		return
	}
	if payload.Email != email {
		t.Errorf("error at validating token's email: %s", err.Error())
		return
	}

}

func TestExpiredPasetoToken(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	r := utils.RandomString(32)
	p, err := NewPasetoFactory(r)
	if err != nil {
		t.Errorf("error at generating object: %s", err.Error())
		return
	}
	email := "admin@email.com"

	token, err := p.GenerateToken(email, -time.Minute)
	if err != nil {
		t.Errorf("error at generating token: %s", err.Error())
		return
	}
	_, err = p.VerifyToken(token)
	if err == nil {
		t.Errorf("should not validate a expired token: %s", err.Error())
		return
	}
}

func TestRandomString(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	r := ""
	i := 0
	for i < 10 {
		r = utils.RandomString(10)
		if len(r) != 10 {
			t.Errorf("unexpected length error")
		}
		i += 1
	}
}
