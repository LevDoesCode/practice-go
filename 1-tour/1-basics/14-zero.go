package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// By default, numer variables contain 0
// bools are false
// strings are empty ""
// %q for a double quoted string representation for strings and slice of bytes
