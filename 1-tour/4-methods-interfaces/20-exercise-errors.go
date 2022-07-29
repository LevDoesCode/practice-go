package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x == 0 {
		return 0, nil
	}
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z0 := 0.0
	z1 := 1.0
	t := 0.0000001

	for math.Abs(z0-z1) > t {
		z0 = z1
		z1 -= (z0*z0 - x) / (2 * z0)
	}
	return z0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(121))
}
