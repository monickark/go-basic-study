package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr [5]int
	arr = [...]int{11, 22, 33, 44, 55}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d : %d \n", i, arr[i])
	}
	fmt.Printf("````````````````````````` \n")

	for i, n := range arr {
		fmt.Printf("%d: %d \n", i, n)
	}
}

func main1() {
	var tmp, sum int
	var str string
	var err error
	// for i := 0; i < 2; i++ {
	// 	fmt.Scanf("%d", &tmp)
	// 	sum = sum + tmp
	// }
	// fmt.Printf("Sum value is %d ", sum)
	// fmt.Println("``````````````````````````````````````");

	for {
		fmt.Scanf("%s \n", &str)
		tmp, err = strconv.Atoi(str)
		if err != nil {
			break
		}
		sum = sum + tmp
	}
	fmt.Printf("The sum is %d ", sum)
}

/**
go run 5.array.go
0 : 11
1 : 22
2 : 33
3 : 44
4 : 55
`````````````````````````
0: 11
1: 22
2: 33
3: 44
4: 55
*/
