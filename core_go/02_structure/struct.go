package structure

import "fmt"

// Person struct with embedded Address struct
type Persons struct {
	Name    string
	Age     int
	Address // Embedded struct
}


