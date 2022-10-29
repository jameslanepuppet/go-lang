package main

import (
	"fmt"
	"math"
)

func sqrt(userInput float64) string { // if statement uses {}
	if userInput < 0 {
		return sqrt(-userInput) + "i"
	}
	return fmt.Sprint(math.Sqrt(userInput)) // uses the defualt operands and returns the resulting string
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}