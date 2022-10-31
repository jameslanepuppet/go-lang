package main

import (
	"fmt"
	"time"
)
// The error type is a built-in interface similar to fmt.Stringer:

type MyError struct {
	When time.Time
	What string
}
// Functions often return an error value, and calling code 
// should handle errors by testing whether the error equals nil
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
// A nil error denotes success; a non-nil error denotes failure.

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}