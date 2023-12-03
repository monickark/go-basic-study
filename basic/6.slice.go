package main

import "fmt"

func main() {
	var arr1 []int
	arr1 = append(arr1, 1, 2, 3, 4)
	printSliceInt(arr1)
	arr11 := arr1[:2]
	arr1[0] = 3
	printSliceInt(arr11)
	fmt.Printf("--------------------")
	arr2 := make([]int, len(arr1))
	printSliceInt(arr1)
	copy(arr2, arr1)
	arr1[0] = 3
	printSliceInt(arr2)
}

func main1() {
	daysOfWeek := [...]string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	weekdays := daysOfWeek[1:6]
	printSlice(daysOfWeek[:])
	daysOfWeek[2] = "Tuesday"
	printSlice(weekdays)
}

func printSlice(slice []string) {
	fmt.Printf("%v : len %d : cap %d", slice, len(slice), cap(slice))
}

func printSliceInt(slice []int) {
	fmt.Printf("%v : len %d : cap %d", slice, len(slice), cap(slice))
}
