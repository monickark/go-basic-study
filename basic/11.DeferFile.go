package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	readAndDeferFile(os.Args[1])
}

func readAndDeferFile(fileName string) {
	var fileBytes [10]byte
	var err error
	var fileSize int
	openedFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	for {
		fileSize, err = openedFile.Read(fileBytes[:])
		defer openedFile.Close()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s  \nsize: %d \n", fileBytes, fileSize)
	}
}

/**
go run 11.DeferFile.go ../commands.txt
go mod ini
size: 10
t helloini
size: 7
*/
