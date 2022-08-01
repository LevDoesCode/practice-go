package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys m:", MapKeys(m))
	m2 := MapKeys[int, string](m)
	fmt.Println("keys m:", m2)
	fmt.Println("keys m:", m2[0])
}

// When invoking generic functions, we can often rely on type inference.
// Note that we don’t have to specify the types for K and V when calling MapKeys,
// the compiler infers them automatically,
// though we could also specify them explicitly (line 16)
