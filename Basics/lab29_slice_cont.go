package main

import "fmt"

func main() {
	names := [4]string{
		"James",
		"Tom",
		"Dick",
		"Harry",
	}
	fmt.Println(names) //[James Tom Dick Harry]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [James Tom] [Tom Dick]

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names) //[James XXX] [XXX Dick]
					   //[James XXX Dick Harry]
}

