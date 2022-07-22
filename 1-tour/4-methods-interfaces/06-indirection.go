package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// We cannot create a method with the same name
// but a different receiver
//
// func (v Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
// }

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)        // The method with pointer receiver is called automatically
	ScaleFunc(&v, 10) // must have a pointer argument

	p := &Vertex{4, 3}
	p.Scale(3)      // As expected, the method with pointer receiver is called
	ScaleFunc(p, 8) // p is already a pointer

	fmt.Println(v, p)
}

// Go interprets v.Scale(2) as (&v).Scale(2) as it automatically
// calls the method with the receiver pointer (Scale)
// Making v.Scale(2) and p.Scale(2) semantically equal
