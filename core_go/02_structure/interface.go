package structure

func init() {
	type Person struct {
		Name string
		Age  int
	}

	// Create an instance of Person
	person := Person{
		Name: "Alice",
		Age:  30,
	}
}	

interface PersonInterface {
	String() string
}

// NewPerson creates a new Person instance
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}