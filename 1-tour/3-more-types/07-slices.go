package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(primes[0], primes[1], primes[2], primes[3])
	fmt.Println(s)
}

// arrays are not resizeable
// slices are more common in workign with array data
// arrays have zero-based indexing, and so do slices
// A range includes the first element but excludes the last one
// also know as a half-open range
