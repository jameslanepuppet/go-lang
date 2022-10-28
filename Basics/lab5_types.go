package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe	bool	   = false
	MaxInt	uint	   = 1<<64 -1
	z		complex128 = cmplx.Sqrt(-5 + 12i)
)
// fmt.Printf = print function
func main() { 
	fmt.Printf("Type: %T Value: %v ", ToBe, ToBe)
// output with just %v = Type: bool Value: false
// output with %v/n = Type: bool Value: false/n
	fmt.Printf("Type: %T Value: %v ", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v ", z, z)
}