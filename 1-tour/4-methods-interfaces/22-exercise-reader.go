package main

import (
	"reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(arr []byte) (int, error) {
	size := len(arr)
	for i := 0; i < size; i++ {
		arr[i] = 'A'
	}
	return size, nil
	// Also
	// arr[0] = 'A'
	// return 1, nil
	// Also
	// for i := range bytes {
	// 	bytes[i] = 65 // Int equivalent for A
	// }
	// return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
