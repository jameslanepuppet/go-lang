package main

import "fmt"

func main(){
	fmt.Println("counting")
// will print out all values until for loop ends
	for userInput := 0; userInput < 10; userInput++ {
		defer fmt.Println(userInput)
	}

	fmt.Println("Complete")
}