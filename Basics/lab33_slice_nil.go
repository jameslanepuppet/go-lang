package main

import "fmt"

func main() {
	var userInput []int
	fmt.Println(userInput, len(userInput), cap(userInput))
	if userInput == nil {
		fmt.Println("Error: Nil!")
	}
}