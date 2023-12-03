package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Centre Point
	radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Centre.X = 10
	w.Circle.Centre.X = 20

	fmt.Printf("%d ", w)
}

/**
go run 10.AnonyField.go
{{{20 0} 0} 0}
*/
