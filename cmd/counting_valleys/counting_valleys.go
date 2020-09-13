package main

func CountingValleys(s []string) int {
	result := 0
	altitude := 0

	for _, v := range s {
		switch v {
		case "D":
			altitude--
		case "U":
			altitude++
		}

		// Check if altitude is descending from sea level
		if v == "D" && altitude == -1 {
			result++
		}
	}
	return result
}
