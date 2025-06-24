package basics

import "fmt"

func loops(){
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
	numbers := []int{1, 2, 3, 4, 5}
	//range loop
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	//while loop equivalent
	//go dont have while loop, but we can alternatively use a for loop to achive the same result
	i := 1
	for i <= 5 {
		fmt.Println("While Loop Iteration:", i)
		i++
	}
	//infinite loop
	// Uncomment the following lines to see an infinite loop in action
	// for {
	// 	fmt.Println("This will run forever!")
	// 	// Add a break condition to exit the loop if needed
	// 	if i > 10 {
	// 		break
	// 	}
	// 	i++
}