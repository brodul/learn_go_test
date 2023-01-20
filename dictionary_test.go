package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("know word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)

	})

	t.Run("unknow word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("banana")

		if err == nil {
			t.Fatal("an error should be returned")
		}
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
