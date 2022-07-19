package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	r := math.Sqrt(x)
	reached := false
	for !reached {
		z -= (z*z - x) / (2 * z)
		if z-r < 0.0000001 {
			reached = true
		}
	}
	return z
}

func main() {
	var x float64 = 15
	fmt.Println("Math solution: ", math.Sqrt(x))
	fmt.Println("My solution: ", sqrt(x))
}
