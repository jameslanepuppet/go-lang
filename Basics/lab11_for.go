package main

import "fmt"

func main() {
	sum := 0
	for userNum := 0; userNum < 10; userNum++ { // ++ increment
		sum += userNum 	// add and assign sum = sum + userNum
	}
	fmt.Println(sum)
}

