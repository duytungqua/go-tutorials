package structure

import "fmt"

// Person struct with embedded Address struct
type Persons struct {
	Name    string
	Age     int
	Addres // Embedded struct
}

// Address struct
type Addres struct {
	City  string
	State string
}

// NewPerson creates a new Person instance with embedded Address
func NewPersons(name string, age int, city string, state string) Persons {
	return Persons{
		Name: name,
		Age:  age,
		Addres: Addres{
			City:  city,
			State: state,
		},
	}
}

func main() {
	person := NewPersons("Alice", 30, "New York", "NY")
	fmt.Println(person) // Output: Name: Alice, Age: 30, City: New York, State: NY
}

