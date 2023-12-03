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

/**
 go run 7.pointers.go
 10 , 0
 10 , 824633761976
 5 , 824633761976
 10 , 824634064928
-------------------------------
 30 , 20
*/
