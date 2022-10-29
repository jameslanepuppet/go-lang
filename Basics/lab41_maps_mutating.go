package main

import "fmt"

func main() {
	userGeo := make(map[string]int)
//The value: 42
	userGeo["Answer"] = 42
	fmt.Println("The value:", userGeo["Answer"])
// The value: 48
	userGeo["Answer"] = 48
	fmt.Println("The value:", userGeo["Answer"])
// The value: 0
	delete(userGeo, "Answer")
	fmt.Println("The value:", userGeo["Answer"])
// The value: 0 Present? false
	v, ok := userGeo["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// NOTE If elem or ok have not yet been declared you could use a short declaration form:

