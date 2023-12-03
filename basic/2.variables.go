package main

import "fmt"

func main() {
	var a int
	a = 5
	fmt.Printf(" | %d", a)

	var b int = 10
	fmt.Printf(" | %d", b)

	c := 20
	fmt.Printf(" | %d", c)

	d, e := 30, 40
	fmt.Printf(" || %d %d", d, e)

}
