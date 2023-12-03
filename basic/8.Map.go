package main

import "fmt"

func main() {
	var input string
	wheelsMap := map[string]int{
		"Car":     10,
		"Auto":    20,
		"Bicycle": 30,
	}
	wheelsMap["Suv"] = 10
	fmt.Println(wheelsMap)
	delete(wheelsMap, "Auto")
	fmt.Println(wheelsMap)

	for Vehicle, Wheels := range wheelsMap {
		fmt.Printf("\n %s, %d \n", Vehicle, Wheels)
	}

	fmt.Scanf("%s", &input)
	wheels, ok := wheelsMap[input]
	if !ok {
		fmt.Printf("Key not found")
	} else {
		fmt.Printf("\n No of wheels : %d ", wheels)
	}
}

/**
go run 8.Map.go
map[Auto:20 Bicycle:30 Car:10 Suv:10]
map[Bicycle:30 Car:10 Suv:10]

 Car, 10

 Bicycle, 30

 Suv, 10
Add
Key not found
**/
