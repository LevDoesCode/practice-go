package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// Formatting   Description                                Verb
// 1.234560e+02 Scientific notation                        %e
// 123.456000   Decimal point, no exponent                 %f
// 123.46       Default width, precision 2                 %.2f
// ␣␣123.46      Width 8, precision 2                       %8.2f
// 123.456      Exponent as needed, necessary digits only  %g
