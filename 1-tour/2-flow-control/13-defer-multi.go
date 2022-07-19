package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// defers are piled onto a LIFO stack
// effectively making the calls in this order
//
// fmt.Println(9) - Last one to be stacked, first one out
// fmt.Println(8)
// ..
// fmt.Println(2)
// fmt.Println(1)
// fmt.Println(0)
