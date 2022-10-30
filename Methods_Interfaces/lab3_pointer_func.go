package main

import (
	"fmt"
	"math"
)

type UserID struct {
	X, Y float64
}
// the Abs and Scale methods rewritten as functions
func Abs(geoLoc UserID) float64 {
	return math.Sqrt(geoLoc.X*geoLoc.X + geoLoc.Y*geoLoc.Y)
}

func Scale(geoLoc *UserID, longitude float64) {
	geoLoc.X = geoLoc.X * longitude
	geoLoc.Y = geoLoc.Y * longitude
}

func main() {
	geoLoc := UserID{3, 4}
	Scale(&geoLoc, 10)
	fmt.Println(Abs(geoLoc))
}