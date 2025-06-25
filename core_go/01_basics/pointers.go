package basics

import "fmt"

func incrementValue(value int) {
	value++
	fmt.Println("Incremented value inside function:", value)
}	

//block comment
/*
	Pointers are variables that store the memory address of another variable.
	Declaring a pointer use the * operator before the type.
	var ptr *int
	To get value of a pointer , we use the & operator to get the address of a variable.
	var value int = 10
	We can not assign a value to a pointer directly, we need to use the address of the variable.
*/

func pointers() {
	var ptr *int // Declare a pointer to an int
	var value int = 10
	ptr = &value

	fmt.Printf("value", value)
	fmt.Printf("Pointer address: %p\n", ptr) // Print the address stored in the
	fmt.Printf("Pointer value: %d\n", *ptr) // Dereference the pointer to get the value


}