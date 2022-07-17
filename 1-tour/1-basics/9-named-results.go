package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	// return statement w/o args returns the named return values (x, y)
	// aka naked return. only used in short functions
}

func main() {
	fmt.Println(split(17))
	fmt.Println(split(13))
}
