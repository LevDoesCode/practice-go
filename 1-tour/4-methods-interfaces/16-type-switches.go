package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case float64:
		fmt.Printf("Thrice %v is %v\n", v, v*3)
	default:
		fmt.Printf("I don't know about type %T", v)
	}
}

func main() {
	do(21)
	do("hellou, deja el show")
	do(3.1416)
	do(true)
}

// A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.
// switch v := i.(type) {
// case T:
//     // here v has type T
// case S:
//     // here v has type S
// default:
//     // no match; here v has the same type as i
// }
// Here we use the type keywords in the switch declaration
