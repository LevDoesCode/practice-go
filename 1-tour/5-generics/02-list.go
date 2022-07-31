package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elements []T
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}

// Using a stringer
func (lst *List[T]) String() string {
	var elements []T
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return fmt.Sprint(elements)
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(24)
	fmt.Println("List:", lst.GetAll())
	fmt.Printf("%T\n", lst)
	fmt.Println("List:", &lst) // We send a pointer to lst
}
