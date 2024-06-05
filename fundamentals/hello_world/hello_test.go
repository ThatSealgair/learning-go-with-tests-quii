package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to the homies", func(t *testing.T) {
		got := Hello("Homies", "English")
		want := "Hello, Homies."

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World.' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World."

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hola when given Spanish", func(t *testing.T) {
		got := Hello("Homies", "Spanish")
		want := "Hola, Homies."

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
