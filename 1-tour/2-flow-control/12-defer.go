package main

import (
	"fmt"
	"time"
)

func deferThis() {
	defer fmt.Println("Seconds before: ", time.Now().Second())
	for i := 0; i < 5000000000; {
		i++
	}
	defer fmt.Println("Seconds after: ", time.Now().Second())
}

func main() {
	deferThis()
	fmt.Println("Seconds now: ", time.Now().Second())
}
