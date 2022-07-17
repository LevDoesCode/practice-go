package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 // this create a new variable with implicit type
	c, python, java := true, false, "No!"

	fmt.Println(i, j, k, c, python, java)
}

// outside functions, however, we can't use the := construct
// only keywords var, func, etc
