package main
// is a dynamically-sized, flexible view into the elements of an array.
// slices are much more common than arrays.
import "fmt"
func main() {
	userID := [6]int{2, 4, 6, 34 ,1, 12}
// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	var userSlice []int = userID[1:4] // output = [4 6 34]
	fmt.Println(userSlice)
}

