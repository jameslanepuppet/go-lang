package main

import (
	"fmt"
	"math"
)
// In go you don't have classes. You define methods or types
// A method is a function with a special reciever argument

type UserID struct {
	X, Y float64
}

type MyFloat float64
// the reciever appears in its own argument list between the func
// keyword and the method name
func (geoLoc UserID) Abs() float64 {
	return math.Sqrt(geoLoc.X*geoLoc.X + geoLoc.Y*geoLoc.Y)
}
// you can declar a method on non-struct typre
// You cannot declare a method with a receiver whose type is defined in another package 
// (which includes the built-in types such as int).
func (f MyFloat) Abs() float64 {
		if f < 0 {
			return float64(-f)
		}
		return float64(f)
}

func main() {
	geoLoc := UserID{3, 4}
	f := MyFloat(-math.Sqrt2)
	fmt.Println(geoLoc.Abs())
	fmt.Println(f.Abs())
}