package main

import "fmt"
// slice has both length and capacity
func main() {
	userInput := []int{2, 3, 4, 5, 6, 7, 2, 3}
	printSlice(userInput) //len = 8 cap = 8

	userInput = userInput[:0]
	printSlice(userInput) //len = 0 cap = 8

	userInput = userInput[:4] //len = 4 cap = 8
	printSlice(userInput)

	userInput = userInput[2:] //len = 2 cap = 8
	printSlice(userInput) 
}

func printSlice(userInput []int) {
	fmt.Printf("len=%d cap=%d $v\n", len(userInput), cap(userInput), userInput)
}