package main

import (
	"fmt"
	"sync"
)

const iterates = 10000

type Counter struct {
	count int
	sync.Mutex
}

func main() {
	num := 0
	counter := Counter{count: 10}
	mutex := sync.Mutex{}
	completed := make(chan struct{})
	go func() {
		for i := 0; i < iterates; i++ {
			mutex.Lock()
			counter.count++
			num++
			mutex.Unlock()
		}
		completed <- struct{}{}
	}()
	go func() {
		for i := 0; i < iterates; i++ {
			mutex.Lock()
			num--
			counter.count--
			mutex.Unlock()
		}
		completed <- struct{}{}
	}()
	<-completed
	<-completed

	fmt.Println(num)
	fmt.Println(counter.count)
}
