package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleNoPointer(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	ScaleNoPointer(v, 10)
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

// Here we changed the methods (bound to a type) to functions
// We see we can send a pointer to Vertex
// We do this by prefixing the ampersand (&) to the variables
// at the time of the function call
