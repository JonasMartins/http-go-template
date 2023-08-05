package test

import (
	"log"
	"os"
	"testing"
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
func TestOneEqualOne(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	got, want := 1, 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
