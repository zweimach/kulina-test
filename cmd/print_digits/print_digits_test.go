package main

import "testing"

func TestPrintDigits(t *testing.T) {
	result := PrintDigits("123456")
	expected := []string{"100000", "20000", "3000", "400", "50", "6"}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("PrintDigits result was incorrect, got: %s, want: %s.", v, expected[i])
		}
	}

	result = PrintDigits("1345679")
	expected = []string{"1000000", "300000", "40000", "5000", "600", "70", "9"}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("PrintDigits result was incorrect, got: %s, want: %s.", v, expected[i])
		}
	}
}
