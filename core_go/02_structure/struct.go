package structure

import "fmt"

// Person struct with embedded Address struct
type Persons struct {
	Name    string
	Age     int
	Address // Embedded struct
}


// NewPerson creates a new Person instance with embedded Address
func NewPerson(name string, age int, city string, state string) Person {
	return Persons{
		Name: name,
		Age:  age,
		Address: Address{
			City:  city,
			State: state,
		},
	}
}