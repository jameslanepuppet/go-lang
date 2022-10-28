package main

import (
	"fmt"
	"math"
)

func multiply(x, n, userValue float64) float64 {
	if num := math.Pow(x, n); num < userValue {
		return num
	}
	return userValue
}

func main() {
	fmt.Println(
		multiply(3, 2, 10),
		multiply(3, 3, 20),
	)
}