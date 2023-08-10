package test

import (
	"log"
	"project/src/pkg/utils"
	"testing"
)

func TestStringMethods(t *testing.T) {
	defer func() {
		log.Println("Deferred TestStringMethods tearing down.")
	}()

	t.Run("Test StringContains Method", func(t *testing.T) {
		str := "test"
		arr := []string{
			"one", "two", "three", "test",
		}

		got := utils.StringContains(&str, &arr)

		want := true
		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}

	})

}
