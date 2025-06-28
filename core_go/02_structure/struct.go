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

//create instance of Person
func CreatePerson() {
	person := Person{
		name: "Alice",
		age:  30,
	}

	fmt.Println(person.String()) // Accessing the String method
	person.age = 35              // Modifying the age field
	fmt.Println(person.String()) // Accessing the String method again	
}

//exported and unexported fields
type PersonWithFields struct {
	Name string // Exported field (accessible outside the package)
	age  int    // Unexported field (not accessible outside the package)
}
func (p PersonWithFields) GetAge() int {
	return p.age // Accessing the unexported field through a method
}	