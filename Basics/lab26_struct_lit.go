package main

import "fmt"

type gpsLocation struct {
	longitude int
	latitude int
}

var (
	gps1 = gpsLocation{1 ,2} //type struct
	gps2 = gpsLocation{longitude: 1} // is implicit
	gps3 = gpsLocation{} //
	userLocation = &gpsLocation{1, 2} // has type *gpsLocation
)

func main() {
	fmt.Println(gps1, userLocation, gps2, gps3)
}