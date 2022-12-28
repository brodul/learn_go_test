package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say Hello to people", func(t *testing.T) {
		got := Hello("brodul")
		want := "Hello, brodul"

		assertCorrectMessage(t, got, want)

	})
	t.Run("say Hello, world if empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
