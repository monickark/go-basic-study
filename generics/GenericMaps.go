package main

import "fmt"

func Keys[K comparable, V any](data map[K]V) []K {
	keys := make([]K, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}
	return keys
}

func main() {
	intMap := map[int]int{
		1: 1,
		2: 4,
	}
	fmt.Println(Keys(intMap))
}
