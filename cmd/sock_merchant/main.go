package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ar := make([]string, 0)
	for i := 0; i < n; i++ {
		var input string
		fmt.Scan(&input)
		ar = append(ar, input)
	}

	fmt.Println(SockMerchant(ar))
}
