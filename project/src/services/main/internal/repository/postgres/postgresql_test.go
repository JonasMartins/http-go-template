package postgres

import (
	"fmt"
	"log"
	"os"
	"project/src/services/main/configs"
	"project/src/services/main/domain/usecases"
	"testing"

	"github.com/gin-gonic/gin"
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
func TestAddUser(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	cfg, err := configs.LoadConfig()
	if err != nil {
		t.Errorf("error on get config %s", err.Error())
		return
	}
	r, err := NewRepository(cfg)
	if err != nil {
		t.Errorf("error on creating repo %s", err.Error())
		return
	}

	data := usecases.AddUserParams{
		Name:     "Test",
		Email:    "test@email.com",
		Password: "test",
	}

	ginEngine := gin.Default()
	c := gin.CreateTestContextOnly(nil, ginEngine)
	response, err := r.AddUser(c, &data)
	if err != nil {
		t.Errorf("error on test createUser %s", err.Error())
		return
	}

	fmt.Println(response)

	got, want := 1, 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
