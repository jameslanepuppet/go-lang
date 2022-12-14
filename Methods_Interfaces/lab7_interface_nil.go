package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}
// If the concrete value inside the interface itself is nil, 
// the method will be called with a nil receiver.
// Note that an interface value that holds a nil concrete value is itself non-nil.
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
