package structure
import "fmt"

func test() {
	type Person struct {
		Name string
		Age  int
	}

	// Create an instance of Person
	person := Person{
		Name: "Alice",
		Age: 30,
	}	
}