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
	//use map to store the elements
	numMap := make(map[int]bool)
	for _, num := range arr {
		complement := target - num
		if _, found := numMap[complement]; found {
			fmt.Printf("Found a pair using map: %d + %d = %d\n", num, complement, target)
		}
		numMap[num] = true
	}
	//use nested loops to find the pairs
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				fmt.Printf("Found a pair using nested loops: %d + %d = %d\n", arr[i], arr[j], target)
			}
		}
	}
}