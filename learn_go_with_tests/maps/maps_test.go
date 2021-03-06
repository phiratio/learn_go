package main

import "testing"

const (
	TestStringKey   = "test"
	TestStringValue = "this is just a test"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{TestStringKey: TestStringValue}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(TestStringKey)
		want := TestStringValue

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := TestStringKey
		definition := TestStringValue

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := TestStringKey
		definition := TestStringValue
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := TestStringKey
		definition := TestStringValue
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		_ = dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := TestStringKey
		definition := TestStringValue
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}
func assertStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
