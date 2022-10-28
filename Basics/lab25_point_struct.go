package main

import "fmt"

type gpsLocation struct {
	longitude int
	latitude int
}
// Struct fields can be accessed through a struct pointer
func main() {
	userPin := gpsLocation{2355, 5768}
	savedPin := &userPin
	savedPin.longitude = 4570 //access the field X of a struct when we have the struct pointer
	fmt.Println(userPin)
}