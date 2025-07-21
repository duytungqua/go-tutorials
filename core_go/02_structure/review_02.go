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