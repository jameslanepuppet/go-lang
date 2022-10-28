package main

import "fmt"
//A pointer holds the memory address of a value.
func main() {
	userInput, userOutput:= 42, 2701
// The type *T is a pointer to a T value. Its zero value is nil
	temp := &userInput 
	fmt.Println(*temp)
	*temp = 21
	fmt.Println(userInput)
// The & operator generates a pointer to its operand.
	temp = &userOutput
	*temp = *temp / 37 //The * operator denotes the pointer's underlying value.
	fmt.Println(userOutput)		
	}