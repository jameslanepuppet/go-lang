package main

import (
	"fmt"
	"time"
)
// Switch without a condition is the same as switch true
// clean way to write long if-then-else chains 
func main() {
	userHour := time.Now()
	switch {
	case userHour.Hour() < 12:
		fmt.Println("Good morning!")
	case userHour.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}