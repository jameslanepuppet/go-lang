package main

import (
	"fmt"
)

func Sqrt(num float64) float64 {
	z := float64(1)
	z -= (z*z - num) / (2*z)
	if z == 10 {
		return z*z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(5))
}