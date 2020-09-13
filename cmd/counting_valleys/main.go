package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, 0)
	for i := 0; i < n; i++ {
		var input string
		fmt.Scan(&input)
		s = append(s, input)
	}

	fmt.Println(CountingValleys(s))
}
