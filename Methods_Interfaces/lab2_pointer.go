package main

import (
	"fmt"
	"math"
)
// You can declare methods with pointer receivers.
type UserID struct {
	X, Y float64
}
//This means the receiver type has the literal syntax *T for some type T. 
// (Also, T cannot itself be a pointer such as *int.)
func (geoLoc UserID) Abs() float64 {
	return math.Sqrt(geoLoc.X*geoLoc.X + geoLoc.Y*geoLoc.Y)
}

func (geoLoc *UserID) Scale(longitude float64) {
	geoLoc.X = geoLoc.X * longitude
	geoLoc.Y = geoLoc.Y * longitude
}

func main() {
	geoLoc := UserID{3, 4}
	geoLoc.Scale(10)
	fmt.Println(geoLoc.Abs())
}