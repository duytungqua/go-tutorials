package structure

//embeddin.go
import "fmt"

// Person struct with embedded Address struct
type Personss struct {
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
func NewPers(name string, age int, city string, state string) Personss {		
	return Personss{
		Name: name,
		Age:  age,
		Address: Address{
			City:  city,
			State: state,
		},
	}
}

func funcs() {
	person := NewPers("Alice", 30, "New York", "NY")
	fmt.Println(person) // Output: {Alice 30 {New York NY}}
	
	// Accessing embedded struct fields
	fmt.Println("City:", person.City)   // Output: City: New York
	fmt.Println("State:", person.State) // Output: State: NY
}