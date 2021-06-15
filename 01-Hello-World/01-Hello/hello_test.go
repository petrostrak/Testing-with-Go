package main

import "testing"

// We write a test for our Hello function
// Run go test in terminal to execute file
func TestHello(t *testing.T) {
	got := Hello("Petros")
	want := "Hello, Petros"

	if got != want {
		// %q	a single-quoted character literal safely escaped with Go syntax.
		t.Errorf("got %q want %q", got, want)
	}
}
