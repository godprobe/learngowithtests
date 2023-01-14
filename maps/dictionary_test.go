package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Query exists", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Non-existent query", func(t *testing.T) {
		_, got := dictionary.Search("404")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	query := "word"
	definition := "this is just a test"

	dictionary.Add(query, definition)

	assertDefinition(t, dictionary, query, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, query, definition string) {
	t.Helper()

	got := d[query]
	want := definition

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
