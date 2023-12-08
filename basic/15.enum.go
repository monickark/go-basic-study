package main

import "fmt"

type DayOfWeek int

const (
	Sunday DayOfWeek = 2 * iota
	Monday
	_
	Tuesday
	Wednesday = 16
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Printf("%d", DayOfWeek(4))
}

/**
go run 15.enum.go
0
2
6
16
16
*/
