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
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		query := "word"
		definition := "this is just a test"

		err := dictionary.Add(query, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, query, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		query := "test"
		initialDefinition := "this is the first entry for test"

		dictionary := Dictionary{
			query: initialDefinition,
		}
		addDefinition := "this is a second entry for test"
		err := dictionary.Add(query, addDefinition)

		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionary, query, initialDefinition)
	})
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

	got, err := d.Search(query)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q, want %q", got, definition)
	}
}
