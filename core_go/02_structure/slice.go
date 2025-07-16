package structure

import "fmt"

/*
	slice is a dynamically-sized, flexible view into the elements of an array
*/

func slice(){
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // Slice from index 1 to 3 (exclusive of 4)
	fmt.Println(slice) // Output: [2 3 4]

	//make slice
	slicem := make([]int, 3, 5)
	
}