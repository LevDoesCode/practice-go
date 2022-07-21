package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// We can declare methods on non-struct types
// Here we have MyFloat which is of float64 type
// Any methods declared need to be in the same package
// as the definition of the type used as receiver
