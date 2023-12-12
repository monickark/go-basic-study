package main

import (
	"fmt"
	p "hello/pkg"
)

func init() {
	fmt.Printf("inside init function \n")
}

func main() {
	p.PrintDetails()
}
