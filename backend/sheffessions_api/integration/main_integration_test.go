package integration

import (
	"testing"
)

func TestIntegrationHelloWorld(t *testing.T) {
	got := "Hello, Integration World!"
	want := "Hello, Integration World!"

	if got != want {
		t.Errorf("Hello, Integration World! = %q, want %q", got, want)
	}
}
