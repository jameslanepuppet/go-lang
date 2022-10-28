package main

import "fmt"

type gpsLocation struct {
	longitude int
	latitude int
}

func main() {
	userLocation := gpsLocation{200, 300}
	userLocation.latitude = 400
	fmt.Println(userLocation.latitude)
}