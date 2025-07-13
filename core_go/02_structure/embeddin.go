package structure

//embeddin.go
import "fmt"

// Person struct with embedded Address struct
type Person struct {
	Name    string
	Age     int
	Address // Embedded struct
}		

// Address struct
type Address struct {
	City  string
	State string
}
// NewPerson creates a new Person instance with embedded Address
func NewPerson(name string, age int, city string, state string) Person {		
	return Person{
		Name: name,
		Age:  age,
		Address: Address{
			City:  city,
			State: state,
		},
	}
}