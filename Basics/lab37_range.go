package main

import "fmt"
// The range form of the for loop iterates over a slice or map.
// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.When ranging over a slice, two values are returned for each iteration. 
// The first is the index 
// the second is a copy of the element at that index.
var geoLoc = []int{1, 2, 3, 4, 5, 6, 7, 8}

// Remember!
// & returns the memory address of the following variable.
//* returns the value of the following variable

func main() {
	for i, v := range geoLoc {
		fmt.Printf("2**%d = %d\n", i, v) //
	}
}