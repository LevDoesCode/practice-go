package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// or also
// x, y := 3, 4
// i := uint(f)
// f := math.Sqrt(float64(x*x + y*y))
