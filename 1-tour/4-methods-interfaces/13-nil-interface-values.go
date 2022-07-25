package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i) // shows <nil>, <nil>
	i.M()       // throws an error because M is not defined for the type i is
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// A nil interface value holds neither value nor concrete type.
// Calling a method on a nil interface is a run-time error
// because there is no type inside the interface tuple to indicate
// which concrete method to call.
