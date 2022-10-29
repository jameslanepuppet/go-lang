package main

import "fmt"

func main() {
	var userOutput [2]string
	userOutput[0] = "Hello"
	userOutput[1] = "world"
	fmt.Println(userOutput[0], userOutput[1])
	fmt.Println(userOutput)

	userID := [6]int{2, 4, 6, 34 ,1, 12}
	fmt.Println(userID)
}