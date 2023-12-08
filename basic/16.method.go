package main

import "fmt"

type Celcius float32
type Farenheit float32

func main() {
	var c Celcius = 36.9
	f := c.toFarenheit()
	fmt.Println(f)
}

func (c Celcius) toFarenheit() Farenheit {
	return Farenheit(9.0/5.0*c + 32)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%.2f%cF", f, 0x00B0)
}

/*
go run 16.method.go
98.42°F
*/
