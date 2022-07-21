package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// We skip values or ommit them because Go will throw errors
// if we don't use values we declare or get from a range
// This throws an error
//
// for i, value := range pow {
// 	fmt.Printf("%d\n", value)
// }
