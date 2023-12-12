package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

type Logger interface {
	Info(string)
}

type ScreenLogger struct{}

type FileLogger struct {
	File *os.File
}

func (ScreenLogger) Info(message string) {
	fmt.Printf("%v: Info: %s ", time.Now(), message)
}

func (l *FileLogger) Info(message string) {
	fmt.Fprintf(l.File, "%v: Info: %s ", time.Now(), message)
}

func NewFileLoggr(fileName string) *FileLogger {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	return &FileLogger{file}
}

func main() {
	var log Logger
	log = &ScreenLogger{}
	log.Info("Hello World")
	fmt.Println(reflect.TypeOf(log))
	fmt.Println("```````````````````````````````")
	log = NewFileLoggr("../commands.txt")
	fmt.Println(reflect.TypeOf(log))
}

/**
go run 18.Interface.go
2023-12-08 18:38:14.3339836 +0530 IST m=+0.006133301: Info: Hello World *main.ScreenLogger
```````````````````````````````
*main.FileLogger
*/
