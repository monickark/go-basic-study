package main

import (
	"fmt"
	"reflect"
)

type Farenheit float64
type Celcius float64
type DefaultType = Celcius

func main() {
	var f Farenheit
	var c Celcius
	var t DefaultType

	f = 10
	c = 30
	t = c
	c = Celcius(f)

	fmt.Printf("%f, %f, %f", f, c, t)
	fmt.Println("\n", reflect.TypeOf(t))
}

/**
go run 14.typeAlias.go
10.000000, 10.000000, 30.000000
 main.Celcius
*/
