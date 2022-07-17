package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Does Go rule?", Truth)
}

// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax.
