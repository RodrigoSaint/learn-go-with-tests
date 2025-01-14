package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (sleeper *SpySleeper) Sleep() {
	sleeper.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := SpySleeper{}

	Countdown(buffer, &sleeper)

	got := buffer.String()
	want := "3\n2\n1\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
