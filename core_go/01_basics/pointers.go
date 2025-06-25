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
	To get value of a pointer, we use the & operator to get the address of a variable.
	var value int = 10
	We can not assign a value to a pointer directly, we need to use the address of the variable.
*/