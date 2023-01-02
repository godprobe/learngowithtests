package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Francesca", "fr")
		want := "Bonjour, Francesca"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Pinocchio", "it")
		want := "Ciao, Pinocchio"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Hans", "de")
		want := "Hallo, Hans"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "es")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
