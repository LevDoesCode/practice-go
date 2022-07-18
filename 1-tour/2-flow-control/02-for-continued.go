package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

// init and post statements are optional
// the contional should evaluate false at some point
// to avoid an infite loop
