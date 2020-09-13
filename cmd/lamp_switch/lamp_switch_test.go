package main

import "testing"

func TestLampSwitch(t *testing.T) {
	result := LampSwitch(10)
	expected := 3

	if result != expected {
		t.Errorf("LampSwitch result was incorrect, got: %d, want: %d.", result, expected)
	}

	result = LampSwitch(100)
	expected = 10

	if result != expected {
		t.Errorf("LampSwitch result was incorrect, got: %d, want: %d.", result, expected)
	}
}
