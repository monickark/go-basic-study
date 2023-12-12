package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Immediate Hi \n")
	go delayPrint()
	fmt.Printf("Ended")
}

func delayPrint() {
	time.Sleep(5 * time.Second)
	fmt.Printf("Delay Hi \n")
}
