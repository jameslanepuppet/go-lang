package main

import (
	"fmt"
	"math"
)
// Both calls to pow return their results before the call to fmt.Println in main begins
func inAndOut(num, userInput, entry float64) float64{
	if value := math.Pow(num, userInput); value < entry { //need to check what math.Pow is?
		return value
	} else {
		fmt.Printf("%g >= %g/n", value, entry)
	}
	//cant use value here though
	return entry
	}

	func main() {
		fmt.Println(
			inAndOut(3, 2, 10),
			inAndOut(3, 3, 20),
				)
	}