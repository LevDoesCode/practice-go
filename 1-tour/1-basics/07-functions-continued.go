package main

import "fmt"

// func add(x int, y int) int { // Original
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Print(add(11, 22))
}
