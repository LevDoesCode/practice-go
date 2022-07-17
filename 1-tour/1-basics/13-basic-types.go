package main

import (
	"fmt"
	"math/cmplx"
)

// this is a variable declaration factored into a block
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	// "1<<64" bit shift left, same as multiplying 1 by 2, 64 times 1*2^64
	// "32>>5" bit shift right, same as dividing 32 by 2, 5 times 32/2^5
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Basic Types
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32
//      // represents a Unicode code point
// float32 float64
// complex64 complex128

// int, uint and uintptr are 32 or 64-bits wide depending on the system
