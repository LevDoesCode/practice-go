package main

import "fmt"

func main() {
	i, j := 42, 500

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	var q *int
	q = &j         // point to j
	*q = *q / 20   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Printf("The address of i: %v\n", &i)
	fmt.Printf("The address stored in p %v\n", p)
	fmt.Printf("The value of i + j is %v\n", *p+*q)
	fmt.Printf("Which is the same as %v\n", i+j)
	fmt.Printf("Value of i: %s\n", fmt.Sprint(i))
	fmt.Printf("Value of j: %s\n", fmt.Sprint(j))
}

// Go doesn't have poiter arithmetic such as
// myPointer + 2
