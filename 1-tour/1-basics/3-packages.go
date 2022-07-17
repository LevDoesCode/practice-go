package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	rand.Seed(int64(time.Now().Second()))
	fmt.Println("My favorite number is", rand.Intn(10))
}
