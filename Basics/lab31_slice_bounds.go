package main

import "fmt"

func main() {
	userID := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(userID)

	userID = userID[1:4] //[3 4 5]
	fmt.Println(userID)

	userID = userID[:2] //[3 4]
	fmt.Println(userID)

	userID = userID[1:] //[4]
	fmt.Println(userID)
}