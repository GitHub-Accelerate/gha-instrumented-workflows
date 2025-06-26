package main

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, world!"
	got := Greet()
	if got != expected {
		t.Errorf("Greet() = %q; want %q", got, expected)
	}
}
