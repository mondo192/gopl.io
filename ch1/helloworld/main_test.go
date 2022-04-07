package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld()
	expected := "Hello, World!"

	if result != expected {
		t.Errorf("Actual: %q, Received: %q", result, expected)
	}
}
