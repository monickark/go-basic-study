package main

import (
	"fmt"
	"time"
)

func main() {
	// create & consume channel
	counter := make(chan int)
	even := make(chan int)
	odd := make(chan int)
	square := make(chan int)
	merge := make(chan int)
	out := make(chan struct{})

	//GoRoutines
	go counterFn(counter)
	go squarerFn(counter, square)
	go counterSplit(counter, even, odd)
	go merger(square, odd, merge)
	go printOut(merge, out)
	<-out
}

func counterFn(counter chan int) {
	for i := 0; i < 5; i++ {
		counter <- i
	}
	close(counter)
}

func squarerFn(counter chan int, square chan int) {
	for i := range counter {
		time.Sleep(1 * time.Second)
		square <- i * i
	}
	close(square)
}

func counterSplit(counter chan int, even chan int, odd chan int) {
	for i := range counter {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func merger(even chan int, odd chan int, merge chan int) {
	i := 0
	for {
		fmt.Printf("%d , ", i)
		a, ok := <-even
		if !ok {
			i++
		} else {
			merge <- a
		}
		a, ok = <-odd
		if !ok {
			i++
		} else {
			merge <- a
		}
		if i >= 3 {
			break
		}
	}
	close(merge)
}

func printOut(merge chan int, out chan struct{}) {
	for i := range merge {
		fmt.Println(i)
	}
	out <- struct{}{}
}

/* go run 25.channel_pipeline.go
0
1
2
3
4
**/
