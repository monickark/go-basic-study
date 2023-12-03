package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%d", &num)
	// if num%2 == 0 {
	// 	fmt.Printf("Entered number %d is even", num)
	// } else {
	// 	fmt.Printf("Entered number %d is odd", num)
	// }
	// fmt.Printf("----------------------------------------")
	// if n := num / 2; num%2 == 0 {
	// 	fmt.Printf("Entered number %d is even & %d times", num, n)
	// } else {
	// 	fmt.Printf("Entered number %d is odd & %d times", num, n)
	// }
	fmt.Printf("%s", fizzbuzzswitch(num))
}

func fizzbuzz(n int) string {
	if n%5 == 0 && n%3 == 0 {
		return "Fizzbuzz"
	} else if n%5 == 0 {
		return "Fizz"
	} else if n%3 == 0 {
		return "buzz"
	}
	return strconv.Itoa(n)
}

func fizzbuzzswitch(n int) string {
	switch {
	case n%5 == 0 && n%3 == 0:
		return "Fizzbuzz"
	case n%5 == 0:
		return "Fizz"
	case n%3 == 0:
		return "buzz"
	}
	return strconv.Itoa(n)
}
