package main

import "fmt"

func main() {
	var c byte
	fmt.Printf("Due to want to subscribe ? Y/N")
	fmt.Scanf("%c", &c)
	switch c {
	case 'Y':
		fmt.Println("Subscribed.....")
	case 'N':
		fmt.Println("Not Subscribed.....")
	}
}

/**
PS D:\Study\Go\GoLearn\basic> go run 4.switch.go
Due to want to subscribe ? Y/NY
Subscribed.....
*/
