package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
	Lel() int64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser. We fix this in line 49
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) Lel() int64 {
	return 1
}

type Vertex struct {
	X, Y float64
}

// To fix line 22, we change (v *Vertex) below to (v Vertex)
// because .Abs(), called in line 24, was defined on *Vertex
// and we cannot declare both. Line 19, we can do this bc v.Abs translates to (&v).Abs
// func (v *Vertex) Abs() float64 {
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f Vertex) Lel() int64 {
	return 1
}

// Line 18, 19, we can change the value of a to f and &v, and v later,
// because both types have implementes ALL (2) methods declared
// in the Abs interface. Otherwise it wouldn't accept and throw an error.
