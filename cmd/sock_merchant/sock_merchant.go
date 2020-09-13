package main

func SockMerchant(ar []string) int {
	result := 0
	sockPairs := make(map[string]int, 0)

	for _, v := range ar {
		if _, ok := sockPairs[v]; ok {
			sockPairs[v]++
		} else {
			sockPairs[v] = 1
		}
	}

	for _, v := range sockPairs {
		result += v / 2
	}
	return result
}
