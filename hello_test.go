package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Test 1: ", func(t *testing.T) {
		got := Hello("Shikhar")
		want := "Hello, Shikhar"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Test 2: ", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
