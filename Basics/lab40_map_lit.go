package main

import "fmt"

type geoLoc struct {
	Lat, Long float64
}

var m = map[string]geoLoc{
	"Bell Labs": geoLoc{
		40.68433, -74.39967,
	},
	"Google": geoLoc{
		37.42202, -122.08408,
	},
}

var n = map[string]geoLoc{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
func main() {
	fmt.Println(m, n)
}
// output map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]