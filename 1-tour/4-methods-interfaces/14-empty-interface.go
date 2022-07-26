package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42 // normally an int
	describe(i)

	// We can assign a different type because i is of interface{} type
	i = 3.141516
	describe(i)

	// Again, we can assign it a string type because it's interface{} type
	i = "hello"
	describe(i)

	var m interface{}
	describe(m)

	// var n *myType
	// m = n
	m = myType{1.5, 3.5}
	fmt.Printf("%v %T\n", m, m)
	// This throws an error because empty interfaces don't implement any methods
	// m.sum(3, 6)
	//m = myInter{}
	var n myInter
	n = &myType{3.34, 34.22}
	describe(n)
	n.sum(3, 9)
	//(&m).sum(3, 6)
}

type myInter interface {
	sum(a, b int) int // A function with 2 variables and a type returned
}

type myType struct {
	x float64
	y float64
}

// The method must take two variables and return an int in order for the
// interface to be satisfied
func (t *myType) sum(a, b int) int {
	fmt.Printf("X + a: %v \nY + b: %v", t.x+float64(a), t.y+float64(b))
	return a + b
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// The interface type that specifies zero methods is known as the empty interface:
// interface{}
// An empty interface may hold values of any type.
// (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type.
// For example, fmt.Print takes any number of arguments of type interface{}.
