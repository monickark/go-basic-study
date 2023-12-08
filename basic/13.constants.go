package main

import "fmt"

const PI float32 = 3.14
const (
	RADIUS = 1.44
	A      = 5
)

func main() {
	var x float32
	var y float64

	x = PI
	y = RADIUS
	fmt.Printf("%f, %f", x, y)
	fmt.Println()
	y = float64(PI)
	fmt.Printf("%f, %f", x, y)

}

/**
 go run 13.constants.go
3.140000, 1.440000
3.140000, 3.140000
*/
