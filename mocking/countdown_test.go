package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const write = "write"
const sleep = "sleep"

type SpyCountdown struct {
	Calls []string
}

func (s *SpyCountdown) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdown) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCounter(t *testing.T) {
	t.Run("test counter", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 but got %d", spySleeper.Calls)
		}
	})

	t.Run("test counter write and sleep order", func(t *testing.T) {
		spyCountdown := &SpyCountdown{}

		Countdown(spyCountdown, spyCountdown)

		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}

		if !reflect.DeepEqual(want, spyCountdown.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spyCountdown.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
