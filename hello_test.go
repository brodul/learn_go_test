package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("brodul")
	want := "Hello, brodul"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
