package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Name string `id:"firstName"`
	Age  int
}

func displaySimple(x interface{}) {
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
}

func displaySwitch(x interface{}) {
	ty := reflect.TypeOf(x)
	val := reflect.ValueOf(x)
	fmt.Println(ty)
	switch val.Kind() {
	case reflect.Struct:
		fmt.Println(ty)
		for i := 0; i < ty.NumField(); i++ {
			fmt.Printf("%s = %v & %s \n", ty.Field(i).Name, val.Field(i), ty.Field(i).Tag.Get("id"))
		}
	default:
		fmt.Println(ty)
		fmt.Println(val)
	}
}

func main() {
	employee := Employee{Name: "Monicka", Age: 10}
	// displaySimple(employee)
	displaySwitch(employee)
}
