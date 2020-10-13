package directory

import (
	"testing"
)

func TestSearch(t *testing.T) {

	//directory := map[string]string{"test": "this is a test."}
	directory := Directory{"test": "this is a test."}

	t.Run("known word", func(t *testing.T) {
		//got := Search(directory, "test")
		got, _ := directory.Search("test")
		want := "this is a test."

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := directory.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		directory := Directory{}
		//directory := make(Directory)

		key := "testAdd"
		value := "this is another test."

		err := directory.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, directory, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is a test."
		directory := Directory{key: value}

		err := directory.Add(key, "existing test.")

		assertError(t, err, ErrKeyExists)
		assertDefinition(t, directory, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is a test."
		directory := Directory{key: value}

		updateValue := "update test."

		err := directory.Update(key, updateValue)

		assertError(t, err, nil)
		assertDefinition(t, directory, key, updateValue)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		updateValue := "this is a test."
		directory := Directory{}

		err := directory.Update(key, updateValue)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Fatalf("got error '%s' want error '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, directory Directory, key string, value string) {
	t.Helper()

	got, err := directory.Search(key)

	if err != nil {
		t.Fatal("got an error but didnt want one")
	}

	assertStrings(t, got ,value)
}
