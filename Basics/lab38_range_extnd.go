package main

import "fmt"

func main() {
	userEntry := make([]int, 10)
	for geoLoc := range userEntry {
		userEntry[geoLoc] = 1 << uint(geoLoc) // == 2**geoLoc
	}
	for _, value := range userEntry {
		fmt.Printf("%d\n", value)
	}
}