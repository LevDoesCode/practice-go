package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// We don't need to declare something similar to
// func (t T) I.M() {}
// This is because it's the only method of name M() with a T receiver
// This would also be implemented implicityly: func (t myType) M() {}
type myType struct {
	x float64
	y float64
}

func (m myType) M() {
	fmt.Println(m.x + m.y)
}

func main() {
	var i I = T{"hello"}
	var j myType = myType{3, 4}
	i.M()
	j.M()
}
