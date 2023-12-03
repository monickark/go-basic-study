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
