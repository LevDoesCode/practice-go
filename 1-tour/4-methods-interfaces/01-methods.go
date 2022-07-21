package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// We associate the Hypotenuse() function to the Vertext struct
// Together with the data, they function as an object of a class
func (v Vertex) Hypotenuse() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	myVertex := Vertex{3, 4} // data
	fmt.Println(myVertex.Hypotenuse())
}
