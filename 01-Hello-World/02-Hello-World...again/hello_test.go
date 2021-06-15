package main

import "testing"

func TestHello(t *testing.T) {
	// Here we are introducing another tool in our testing arsenal, subtests. Sometimes it is useful to group tests
	// around a "thing" and then have subtests describing different scenarios
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Petros")
		want := "Hello, Petros"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty strin is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
