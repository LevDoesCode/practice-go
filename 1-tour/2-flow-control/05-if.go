package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x) + 1)
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	a, b, c := 4, 5, 6
	fmt.Println(fmt.Sprint(a, b, c))
	var x, y, z string = "eks", "why", "zee"
	fmt.Println(fmt.Sprint(x, y, z))
}

// Sprint can convert a number to a string
// Sprint adds a space when non-string operands
// are used, but doesn't when string operads are used
