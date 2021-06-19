package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	GreetToWriter(&buffer, "Petros")

	got := buffer.String()
	want := "Hello, Petros"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}
