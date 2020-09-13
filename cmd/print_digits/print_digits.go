package main

import (
	"fmt"
	"strings"
)

func PrintDigits(d string) []string {
	result := make([]string, 0, len(d))

	for i, n := range d {
		result = append(result, fmt.Sprintf("%s%s", string(n), strings.Repeat("0", len(d)-i-1)))
	}
	return result
}
