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
