package main

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

func declarePointer(){
	var pointer *int
	var value int = 42
	pointer = &value // Should be assign address of a variable, not a value
	fmt.Printf("Pointer address: %p\n", pointer) // Print the address stored in the pointer

}

func sliceExample(){
	arr := []int{1,2,3,4,5}
	slice := make([]int, 3, 5)

	slice = append(slice, 6, 7)
	fmt.Println("Slice after append:", slice)
	newSlice := slice[1:4]
	fmt.Println("New slice:", newSlice) // Output: [2 3 4]
}


func main(){
	sliceExample()
}