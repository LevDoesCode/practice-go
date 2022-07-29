package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!") // Returns a *strings.Reader
	fmt.Printf("Type: %T\n", r)
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// The bytes that are not overwritten stay, in this case, 32 and 82
		if err == io.EOF {
			break
		}
	}
}
