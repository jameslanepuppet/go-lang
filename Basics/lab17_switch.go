package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("go runs on")
	switch os := runtime.GOOS; os {
	case "test":
			fmt.Println("OS X.")
	case "Linux":
			fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

