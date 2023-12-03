package main

import "fmt"

func main() {
	var a int
	var b *int
	a = 10
	fmt.Printf(" %d , %d ", a, b)
	b = &a
	fmt.Printf("\n %d , %d ", a, b)
	a = 5
	fmt.Printf("\n %d , %d ", a, b)
	*b = 10
	fmt.Printf("\n %d , %d ", a, &b)
	fmt.Printf("\n-------------------------------")

	c, d := 20, 30
	swap(&c, &d)
	fmt.Printf("\n %d , %d ", c, d)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
