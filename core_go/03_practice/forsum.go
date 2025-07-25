package practice

import "fmt"	

func Sum() {
	var sum int = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of first 10 natural numbers is:", sum)
	
	var numbers = []int{1, 2, 3, 4, 5}
	var total int = 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println("Sum of the slice is:", total)
}
s