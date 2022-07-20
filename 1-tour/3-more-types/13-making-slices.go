package main

import "fmt"

func main() {
	a := make([]int, 5) // slice of len 5, cap 5
	fmt.Println(cap(a))
	printSlice("a", a)

	b := make([]int, 0, 7) // slice of len 0, cap 7
	fmt.Println(cap(b))
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// this will throw an error because the capacity
	// of the underlying array is just 5
	// e := a[0:10]
	// printSlice("e", e)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
