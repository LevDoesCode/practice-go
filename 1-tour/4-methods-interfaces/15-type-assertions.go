package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}

// We asset the type that is in the underlying concrete value by doing this:
// t := myInterface.(myType)
// OR
// t := i.(float64)
// T receives a value for the underlying value
// If we add another variable, it will hold the assetion result
// val, assert := i.(float64)
// The value is 0 if false and the assertion is true or false
