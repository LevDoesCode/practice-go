package main

import (
	"fmt"
	"time"
)

func someFunc() time.Weekday {
	return time.Wednesday
}

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("Today is: ", today)
	fmt.Println("Saturday is: ", time.Saturday)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case someFunc():
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}
}

// case someFunc won't be called if any of the previous cases evaluate to true
// nor will any other cases evaluated, including the default
