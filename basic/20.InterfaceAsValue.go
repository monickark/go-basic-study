package main

import "fmt"

type Employee struct {
	Name string
}

type Manager struct {
	Title string
}

type HashTitle interface {
	GetTitle() string
}

func (m Manager) GetTitle() string {
	return m.Title
}

func main() {
	var data interface{}
	data = int(32)
	data = "This"
	fmt.Println(data)
	x := "Monicka"
	print(x)
	emp := Employee{Name: "Akilan"}
	print(emp)
	manager := Manager{Title: "Associate manager"}
	print(manager)
}

func print(data interface{}) {
	if val, ok := data.(string); ok {
		fmt.Printf("%s\n", val)
	} else {
		fmt.Println("Unsupported Type \n")
	}

	switch val := data.(type) {
	case string:
		fmt.Printf("%s is a string \n", val)
	case int:
		fmt.Printf("%d is an integer \n", val)
	case Employee:
		fmt.Printf("%s is an Employee \n", val.Name)
	case HashTitle:
		fmt.Printf("%s has title of an manager", val.GetTitle())
	default:
		fmt.Printf("Unsupported type", val)
	}
}

/**
go run 20.InterfaceAsvalue.go
This
Monicka
Monicka is a string
------------------------
go run 20.InterfaceAsvalue.go
This
Unsupported Type

123 is an integer
*/
