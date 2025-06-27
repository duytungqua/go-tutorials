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

type Person struct {
	name string
	age  int
}	

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

// NewPerson creates a new Person instance
func NewPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}