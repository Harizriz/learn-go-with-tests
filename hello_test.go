package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Monsiour", "French")
		want := "Bonjour, Monsiour"
		assertCorrectMessage(t, got, want)
	})
}

// testing.TB satisfies both testing.T and testing.B interface
// (got string, want string) is shorten to (got, want string) since it's the same type
// Refactor and improve readability
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // when this func fails, the line number reported will in the function call rather than the test helper
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}