package basics

import "fmt"

func dataTypes() {

	var age int = 30
	var name string = "D"
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	var scores = []int{10, 20, 30, 40, 50}
	fmt.Println("Scores:", scores)
	var isActive bool = true // boolean type
	fmt.Println("Is Active:", isActive)

	person := struct {
		Name string
		Age  int
	}{"D", 30}
	fmt.Println("Person:", person)
	fmt.Println("Person Name:", person.Name)
}