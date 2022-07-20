package main

import "fmt"

func main() {
	var s []int // This creates a slice, not an array
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
