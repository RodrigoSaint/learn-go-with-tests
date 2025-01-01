package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Rodrigo", "")
		want := "Hello Rodrigo"

		assertMessage(t, got, want)
	})
	t.Run("says generic hello", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertMessage(t, got, want)
	})

	t.Run("says hello in Spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola World"

		assertMessage(t, got, want)
	})

	t.Run("says hello in French", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour World"

		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
