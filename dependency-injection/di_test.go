package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("testing greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Shikhar")

		got := buffer.String()
		want := "Hello! Shikhar"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
