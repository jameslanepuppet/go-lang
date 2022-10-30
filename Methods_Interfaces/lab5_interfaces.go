package main

import "fmt"

type UserInter interface {
	M()
}

type UserStruct struct {
	S string
}
// This method means type userStruct implements the interface
// but we dont need to explicitl declare that it does so
func (t UserStruct) M() {
	fmt.Println(t.S)
}
func main() {
	var i UserInter = UserStruct{"Hello"}
	i.M()
}