package main

import "testing"

func TestDict(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("affect")

		assertError(t, err, ErrNotFound)
	})

	t.Run("add word", func(t *testing.T) {
		word := "passion"
		definition := "strong desire"
		dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "passion"
		definition := "affection"
		dict := Dictionary{word: definition}
		err := dict.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("update word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := Dictionary{word: definition}
		newDefinition := "new definition"
		dict.Update(word, newDefinition)

		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("non existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := Dictionary{}
		err := dict.Update(word, definition)

		assertError(t, err, ErrWordNotExists)
	})

	t.Run("delete word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := Dictionary{word: definition}

		dict.Delete(word)

		_, err := dict.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertString(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %q, want %q", got, want)
	}
}

func assertError(tb testing.TB, got, want error) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(tb testing.TB, dict Dictionary, word, definition string) {
	tb.Helper()

	got, err := dict.Search(word)
	if err != nil {
		tb.Fatal("should find added word:", err)
	}
	assertString(tb, got, definition)
}
