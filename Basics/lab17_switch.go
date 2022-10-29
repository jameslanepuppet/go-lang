package main

import (
	"fmt"
	"runtime"
)
// shorter way to write a sequence of if - else statements
func main() {
	fmt.Println("go runs on")
	switch os := runtime.GOOS; os {
	case "test":
			fmt.Println("OS X.") //break statement required at end of each case
	case "Linux":
			fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

