package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Normally, only a copy of the variable is sent to the method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Using a pointer allows access to the origional variable
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleNoChange(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.ScaleNoChange(10)
	fmt.Println(v) // We see now change made
	v.Scale(10)
	fmt.Println(v) // We can see the change to the vertex
	fmt.Println(v.Abs())
}
