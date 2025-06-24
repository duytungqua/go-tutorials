package basics

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}
func add(a int, b int) int {
	return a + b
}
func multiply(a int, b int) int {
	return a * b
}

func resMultivalue() (int, int) {
	a := 10
	b := 20
	return a, b
}	

func rectangleArea(length, width float64) float64 {
	return length * width
}	