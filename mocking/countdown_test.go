package main

import (
	"bytes"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("test counter", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer)

		got := buffer.String()
		want := "3"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
