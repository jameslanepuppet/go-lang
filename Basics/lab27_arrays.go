package main

import "fmt"

func main() {
	var userOutput [2]string
	userOutput[0] = "Hello"
	userOutput[1] = "world"
	fmt.Println(userOutput[0], userOutput[1]) // Hello world
	fmt.Println(userOutput) // [Hello world]

	userID := [6]int{2, 4, 6, 34 ,1, 12} // [2 4 6 34 1 12]
	fmt.Println(userID)
}