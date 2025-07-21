package structure

import "fmt"

//structs
type Employee struct {
	Id int
	Name string
	Phone string
}

//embedding employee struct in department struct
type Department struct {
	Id int
	Name string
	Employee // Embedded struct
}

//value receiver method 
/*
	func (v T) method(){
	}

	1. Encapsulates we dont need to pass the struct as a parameter when calling the method
	2. Can access the fields of the struct directly
	3. Immutable, cannot modify the originnal struct
*/

func (d Department) GetEmployeeName() string {
	return d.Employee.Name
}

