package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	r := s

	fmt.Printf("Type of s: %T\n", s)
	s = s[1:4]
	fmt.Println(s)
	fmt.Printf("Type of s: %T\n", s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	fmt.Println(r)
}

// Similar to as in Python, these slice expressions are equivalent
// for an array var a [10]int
//
// a[0:10]
// a[:10]
// a[0:]
// a[:]
