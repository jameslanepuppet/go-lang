package main 

import "fmt"
// Slices can be created with the built-in make function;
// this is how you create dynamically-sized arrays.
// The make function allocates a zeroed array and 
// returns a slice that refers to that array:
func main() {
	a := make([]int, 5)
	printSlice("a", a) // a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) // d len=3 cap=3 [0 0 0]
}

func printSlice(userID string, x []int) {
	fmt.Printf("%userID len=%d cap=%d %v\n", userID, len(x), cap(x), x)
} 