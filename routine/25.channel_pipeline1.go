package main

import (
	"fmt"
)

func main() {
	// create & consume channel
	in := make(chan int)
	process := make(chan int)
	out := make(chan int)

	go counter(in)
	go squarer(in, process)
	go printer(process, out)
	for i := range out {
		fmt.Println(i)
	}
}

func counter(in chan int) {
	for i := 0; i < 5; i++ {
		in <- i
	}
	close(in)
}

func squarer(in chan int, process chan int) {
	for i := range in {
		process <- i * i
	}
	close(process)
}

func printer(process chan int, out chan int) {
	for i := range process {
		out <- i
	}
	close(out)
}

/*
go run 25.channel_pipeline.go
0
1
2
3
4
**/
