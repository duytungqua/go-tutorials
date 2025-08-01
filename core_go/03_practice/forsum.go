package practice

import "fmt"	

func F4Sum() {
	
	//use two pointers to iterate through the array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 10
	for i, j := 0, len(arr)-1; i < j; {
		sum := arr[i] + arr[j]
		if sum == target {
			fmt.Printf("Found a pair: %d + %d = %d\n", arr[i], arr[j], target)
			i++
			j--
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

}