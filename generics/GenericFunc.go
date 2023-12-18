package main

import "fmt"

type Generic interface {
	int | float32 | float64
}

// GENERIC METHODS
func mathGeneric[T int | float64](x, y T) {
	fmt.Printf("generic output : %v \n", x*y)
}

// GENERIC METHODS
func mathGenInter[T Generic](x, y T) {
	fmt.Printf("generic output : %v \n", x*y)
}

// OLD METHOS WITHOUT GENERICS
func mathInt(num1, num2 int) int {
	return num1 * num2
}

func mathFloat32(num1, num2 float32) float32 {
	return num1 * num2
}

func main() {
	// old methods
	fmt.Printf("Int sum :  %d \n", mathInt(5, 2))
	fmt.Printf("Int float32 : %f \n", mathFloat32(2.5, 5.5))

	// Generic Methods
	mathGeneric(5, 2)
	mathGeneric(2.1, 3.2)

	// Generis methods using interface
	mathGenInter(5, 2)
	mathGenInter(2.1, 3.2)
}
