package main

import "fmt"

type gpsLocation struct {
	longitude int
	latitude int
}

func main() {
	userLocation := gpsLocation{200, 300}
	userLocation.latitude = 400 // Struct fields are accessed using a dot.
	fmt.Println(userLocation.latitude)
}