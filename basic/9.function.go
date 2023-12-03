package main

import "fmt"

func main() {
	x, y := div(55, 3)
	fmt.Printf("%d : %d \n", x, y)

	fn := func(a int) int {
		return a * a
	}
	fmt.Println(fn(2))
	fmt.Println("-----------------------------")
	result := sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("%d", result)

}

func div(num int, denom int) (int, int) {
	return num / denom, num % denom
}

func sum(n ...int) int {
	result := 0
	for _, val := range n {
		result += val
	}
	return result
}
