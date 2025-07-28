package practice

import "fmt"	

func Sum() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", slice)

	// Append elements to the slice
	slice = append(slice, 6, 7)
	fmt.Println("Slice after append:", slice)

	// Remove the last element
	slice = slice[:len(slice)-1]
	fmt.Println("Slice after removing last element:", slice)

	// Length and capacity of the slice
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))
}