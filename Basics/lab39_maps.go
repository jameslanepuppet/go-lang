package main

import "fmt"
// A map maps keys to values
// A zero value of a map is nil. A nil map has no keys, nor can keys be added
type GeoLoc struct {
	Latitude, Longitude float64
}

var userLoc map[string]GeoLoc

func main() {
	userLoc = make(map[string]GeoLoc) // The make function returns a map of the given type ready to use
	userLoc["Bell labs"] = GeoLoc{
		40.456, -73.43232,
	}
	fmt.Println(userLoc["Bell labs"])
}
// output {40.456 -73.43232}
