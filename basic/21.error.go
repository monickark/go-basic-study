package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	text, err := readFile("../commands.txt")
	if err != nil {
		fmt.Printf("err: %s", err)
	} else {
		fmt.Println(text)
	}

}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
