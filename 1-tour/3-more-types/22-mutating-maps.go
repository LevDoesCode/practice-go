package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // we evaluate if the key is present
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 99
	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// v (key) defaults to 0 if not present
// ok (exists) defaults to false if not present
