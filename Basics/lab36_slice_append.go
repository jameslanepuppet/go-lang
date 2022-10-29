package main

import "fmt"
// The resulting value of append is a slice containing all the elements 
// of the original slice plus the provided values.

func main() {
	// Append works on nil slices len=0 cap=0 []
	var userID []int
	printSlice(userID)

// The slice grows as needed len=1 cap=1 [0]
	userID = append(userID, 0)
	printSlice(userID)

// We can add more than one element at a time len=2 cap=2 [0 1]
	userID = append(userID, 1)
	printSlice(userID)

// len=5 cap=6 [0 1 2 3 4]
	userID = append(userID, 2, 3, 4)
	printSlice(userID)
}
func printSlice(userID []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(userID), cap(userID), userID)
}