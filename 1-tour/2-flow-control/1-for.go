package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // init statement, condition expression, post statement
		sum += i
	}
	fmt.Println(sum)
}

// the {} are required in the for
// no parenthesis are required for the compoments
