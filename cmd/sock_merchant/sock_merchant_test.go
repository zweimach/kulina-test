package main

import "testing"

func TestSockMerchant(t *testing.T) {
	result := SockMerchant([]string{"10", "20", "20", "10", "10", "30", "50", "10", "20"})
	expected := 3

	if result != expected {
		t.Errorf("SockMerchant result was incorrect, got: %d, want: %d.", result, expected)
	}

	result = SockMerchant([]string{"1", "2", "1", "2", "1", "3", "2"})
	expected = 2

	if result != expected {
		t.Errorf("SockMerchant result was incorrect, got: %d, want: %d.", result, expected)
	}
}
