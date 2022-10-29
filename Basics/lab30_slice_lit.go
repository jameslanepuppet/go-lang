package main

import "fmt"

func main() {
	userID := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(userID)

	userPaid := []bool{true, false, true, false, true}
	fmt.Println(userPaid)

	userOwed := []struct {
		amount int
		over30Days bool
	}{
		{2, true},
		{3, false},
		{4, true},

	}
	fmt.Println(userOwed)
}