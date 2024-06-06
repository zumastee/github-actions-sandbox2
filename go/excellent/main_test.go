package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even" {
		t.errof("expected: event, actual: %s", result)
	}
}