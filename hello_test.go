package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("waffles")
		want := "Hello, waffles"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("say hello with empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
