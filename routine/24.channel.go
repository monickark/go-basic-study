package main

import (
	"fmt"
)

func main() {
	// create & consume channel
	ch := make(chan int)
	go delayPrintInt(1, ch)
	sum := <-ch
	fmt.Println("Sum is : ", sum)

	// channel overflow
	ch1 := make(chan int, 2)
	go overloadChannel(ch1)
	<-ch1

	//channel generator
	ch2 := make(chan int)
	go ChannelGen(ch2)
	for num := range ch2 {
		fmt.Println(num)
	}
}

func delayPrintInt(num int, ch chan int) {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	ch <- sum
}

func ChannelGen(ch chan int) {
	for i := 1000; i < 1005; i++ {
		ch <- i
	}
	close(ch)
}

func overloadChannel(ch1 chan int) {
	ch1 <- 1
	ch1 <- 2
	fmt.Printf("overload channel val: %d ", <-ch1)
}

/*
go run 24.channel.go
Sum is :  10
overload channel val: 2 1000
1001
1002
1003
1004
**/
