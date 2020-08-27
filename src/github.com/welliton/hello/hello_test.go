package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Welliton")
	want := "Hello, Welliton"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
