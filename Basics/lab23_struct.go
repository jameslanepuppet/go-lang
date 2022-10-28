package main

import "fmt"

type userInput struct {
	latitude int
	longitude int
}

func main() {
	fmt.Println(userInput{1, 2})
}