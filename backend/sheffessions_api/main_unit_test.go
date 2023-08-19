package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := "Hello, Unit World!"
	want := "Hello, Unit World!"

	if got != want {
		t.Errorf("Hello, Unit World() = %q, want %q", got, want)
	}
}
