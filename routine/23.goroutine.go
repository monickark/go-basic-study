package main

import (
	"fmt"
	"time"
)

func main() {
	// delay print
	delayPrintStr("delay print")

	// go routine d
	go delayPrintStr("First")
	go delayPrintStr("Second")
	time.Sleep(1 * time.Second)
}

func delayPrintStr(str string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s %d \n", str, i)
		time.Sleep(100 * time.Millisecond)
	}
}

/*
go run 23.goroutine.go
Third, 0
First, 0
Second, 0
Third, 1
Second, 1
First, 1
First, 2
Second, 2
Third, 2
Second, 3
Third, 3
First, 3
First, 4
Second, 4
Third, 4
1, 0
1, 1
1, 2
1, 3
1, 4
Sum is :  10
**/
