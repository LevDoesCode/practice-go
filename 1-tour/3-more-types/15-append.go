package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	//
	arr := [5]int{1, 2, 3, 4, 5}
	d := arr[:2]
	printSlice(d)

	d = append(d, 8) // We modify the underlying array
	fmt.Println(arr)
	printSlice(d)

	arr = [5]int{1, 2, 3, 4, 5}
	d = arr[:]
	printSlice(d)

	// Here a new array is created different than arr
	d = append(d, 6)
	fmt.Println(arr)
	printSlice(d)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// When we append to a slice, it adds or modifies the value of the underlying array
// If the array is too short, a new one is created and the values
// Are copied over
