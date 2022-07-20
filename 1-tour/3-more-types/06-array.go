package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	booleans := []bool{true, false, true} // This is actually a slice :)

	strings := [...]string{"Penn", "Teller"}

	fmt.Println(primes)
	fmt.Println(booleans)
	fmt.Println(strings)
	fmt.Println(len(primes))
	fmt.Println(len(booleans))
	fmt.Println(len(strings))

	fmt.Printf("Type of primes: %T\n", primes)
	fmt.Printf("Type of booleans: %T\n", booleans)
	fmt.Printf("Type of strings: %T\n", strings)
}
