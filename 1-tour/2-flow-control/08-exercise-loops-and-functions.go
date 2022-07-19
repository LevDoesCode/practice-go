package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	r := math.Sqrt(x)

	for math.Abs(z-r) > 0.0000001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	var x float64 = 15
	fmt.Println("Math solution: ", math.Sqrt(x))
	fmt.Println("My solution: ", sqrt(x))
}
