package structure

//methods
import "fmt"

type Person struct {
	Name string // Exported field (accessible outside the package)
	age  int    // Unexported field (not accessible outside the package)
}	
func (p Person) String() string { // Value receiver
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.age)
}	
// NewPerson creates a new Person instance
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		age:  age,
	}
}	
//create instance of Person
func CreatePerson() {
	person := Person{
		Name: "Alice",
		age:  30,
	}

	fmt.Println(person.String()) // Accessing the String method
	person.age = 35              // Modifying the age field
	fmt.Println(person.String()) // Accessing the String method again	
}