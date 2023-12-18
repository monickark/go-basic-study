package main

import (
	"fmt"
	"sync"
	"time"
)

var config map[string]string
var runOnce sync.Once = sync.Once{}

func loadConfig() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("loading configuration")
	for i := 0; i < 5; i++ {
		config = map[string]string{
			"domain": "localhost",
		}
	}
}

func getConfig(key string) string {
	if config == nil {
		runOnce.Do(loadConfig)
	}
	return config[key]
}

func doSomething(completed chan struct{}) {
	getConfig("domain")
	completed <- struct{}{}
}

func main() {
	completed := make(chan struct{})
	for i := 0; i < 5; i++ {
		go doSomething(completed)
	}
	for i := 0; i < 5; i++ {
		<-completed
	}
}
