package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

// testing.TB accecpts both testing.T and testing.B
// testing.TB is an interface
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // helps debug errors
	if got != want {
		t.Errorf("got %q, and want %q", got, want)
	}
}
