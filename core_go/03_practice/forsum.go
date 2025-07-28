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

	// Iterate over the slice using a for loop
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Element at index %d: %d\n", i, slice[i])
	}
	// Iterate over the slice using a range loop
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}