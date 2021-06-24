package main

import "testing"

func TestHello(t *testing.T) {
	// For helper functions, it's a good idea to accept a testing.TB which is an interface
	// that *testing.T and *testing.B both satisfy.
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper. By doing
		// this when it fails the line number reported will be in our function call rather than
		// inside our test helper.
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	// Here we are introducing another tool in our testing arsenal, subtests. Sometimes it is useful to group tests
	// around a "thing" and then have subtests describing different scenarios
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Petros", "spanish")
		want := "Hola, Petros"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty strin is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
