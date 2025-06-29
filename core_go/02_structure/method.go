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