package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Test 1: ", func(t *testing.T) {
		got := Hello("Shikhar")
		want := "Hello, Shikhar"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test 2: ", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
