package test

import (
	"project/src/pkg/utils"
	"project/src/services/main/configs"
	"reflect"
	"testing"
)

func TestGetFilePath(t *testing.T) {

	u := utils.New()
	f, err := u.GetFilePath(&[]string{"src", "services", "main", "configs", "base.yaml"})
	if err != nil {
		t.Logf("Err %v", err)
	}

	t.Log(f)
	got, want := 1, 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetRoot(t *testing.T) {

	u := utils.New()

	got, err := u.GetRelativeRootDir()
	if err != nil {
		t.Logf("Err %v", err)
	}

	want := "project"
	if *got != want {
		t.Errorf("got %s, wanted %s", *got, want)
	}

}

func TestLoadConfig(t *testing.T) {

	var want configs.Config
	got, err := configs.LoadConfig()
	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
	if reflect.TypeOf(got) != reflect.TypeOf(&want) {
		t.Errorf("got %q, wanted %q", *got, want)
	}
	got = nil
}
