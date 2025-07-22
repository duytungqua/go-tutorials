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
