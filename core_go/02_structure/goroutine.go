package structure

import "fmt"

//goroutine is a lightweight thread of execution managed by the Go runtime, it allows you to run functions concurrently.

/*
	Features of Goroutines:
	1. Lightweight: Gorotines are much lighter than threads, allowing you to run thousands without signigicant overhead.
	2. Concurrent: They can run concurrently, making it easy to write programs
	3. Start with the `go` keyword followed by a function call.
*/
func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}

func routine(){
	fmt.Println("Main function is running...")

	for i := 1; i <= 5; i++ {
		go printNumbers()
	}

	fmt.Println("Main function finished.")
	
	// Wait for goroutines to finish
	var input string
	fmt.Scanln(&input) // Wait for user input to prevent the program from exiting immediately
	fmt.Println("Exiting program.")
}