package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var prev1, prev2 int
	prev1 = 0
	prev2 = 1
	return func() int {
		//fmt.Println("1 is: ", prev1, " 2 is: ", prev2)
		if prev1 == 0 {
			prev1 = 1
			return 0
		}
		if prev1 == 1 && prev2 == 1 {
			prev2 = 2
			return 1
		}

		temp1 := prev1
		temp2 := prev2
		prev2 = prev1 + prev2
		prev1 = temp2
		return temp1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
