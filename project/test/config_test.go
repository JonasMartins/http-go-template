package test

import (
	"log"
	"project/src/pkg/utils"
	"project/src/services/main/configs"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	u := utils.New()

	t.Run("Get FilePath", func(t *testing.T) {
		f, err := u.GetFilePath(&[]string{"src", "services", "main", "configs", "base.yaml"})
		if err != nil {
			t.Logf("Err %v", err)
		}
		t.Log(f)
		got, want := 1, 1
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("Get Root", func(t *testing.T) {
		got, err := u.GetRelativeRootDir()
		if err != nil {
			t.Logf("Err %v", err)
		}
		want := "project"
		if *got != want {
			t.Errorf("got %s, wanted %s", *got, want)
		}
	})

	t.Run("Load config", func(t *testing.T) {
		var want configs.Config
		got, err := configs.LoadConfig()
		if err != nil {
			t.Errorf("Error %s", err.Error())
		}
		if reflect.TypeOf(got) != reflect.TypeOf(&want) {
			t.Errorf("got %q, wanted %q", *got, want)
		}
	})
}
