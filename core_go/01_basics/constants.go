package basics

import (
	"fmt"
	"testing"
)
const (
	PI       = 3.14
	E        = 2.718
	GOLDEN   = 1.6180339887
	MAX_SIZE = 100
)
func TestConstants(t *testing.T) {			
	fmt.Println("Constants:")
	fmt.Println("PI:", PI)
	fmt.Println("E:", E)
	fmt.Println("GOLDEN:", GOLDEN)
	fmt.Println("MAX_SIZE:", MAX_SIZE)
	t.Logf("Constants: PI=%v, E=%v, GOLDEN=%v, MAX_SIZE=%v", PI, E, GOLDEN, MAX_SIZE)
}		