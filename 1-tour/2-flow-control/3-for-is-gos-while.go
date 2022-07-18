package main

import "fmt"

func main() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

// We can drop the semi colons and
// for behaves exactly as a C while statement
// The semi colons also get dropped by auto-formatting
