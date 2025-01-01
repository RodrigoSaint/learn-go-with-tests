package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Rodrigo", English)
		want := "Hello Rodrigo"

		assertMessage(t, got, want)
	})
	t.Run("says generic hello", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello World"

		assertMessage(t, got, want)
	})

	t.Run("says default hello in english", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertMessage(t, got, want)
	})

	t.Run("says hola in Spanish", func(t *testing.T) {
		got := Hello("", Spanish)
		want := "Hola World"

		assertMessage(t, got, want)
	})

	t.Run("says Bonjour in French", func(t *testing.T) {
		got := Hello("", French)
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
