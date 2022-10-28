package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	const name, dept = "James", "CS"

	s := fmt.Sprint(name, " is a ", dept, "Portal.\n")

	io.WriteString(os.Stdout, s)
}