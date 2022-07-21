package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Point struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

// Methods are functions with a receiver (struct) argument
// Above, Abs is a regular function
//
// Methods can only have one receiver
// trying to add two would throw an error
// func (v Vertex, p Point) Abs() float64 {
// }
