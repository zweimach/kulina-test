package main

import "math"

func LampSwitch(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		// Check if i is a square number
		if i == int(math.Pow(math.Round(math.Sqrt(float64(i))), 2)) {
			result++
		}
	}
	return result
}
