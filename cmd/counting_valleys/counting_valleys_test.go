package main

import "testing"

func TestCountingValleys(t *testing.T) {
	result := CountingValleys([]string{"U", "D", "D", "D", "U", "D", "U", "U"})
	expected := 1

	if result != expected {
		t.Errorf("CountingValleys result was incorrect, got: %d, want: %d.", result, expected)
	}

	result = CountingValleys([]string{"D", "D", "U", "U", "U", "U", "D", "D"})
	expected = 1

	if result != expected {
		t.Errorf("CountingValleys result was incorrect, got: %d, want: %d.", result, expected)
	}
}
