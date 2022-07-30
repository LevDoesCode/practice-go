package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(arr []byte) (int, error) {
	n, e := rot.r.Read(arr)
	for i := 0; i < n; i++ {
		if (arr[i] >= 'A' && arr[i] < 'N') || (arr[i] >= 'a' && arr[i] < 'n') {
			arr[i] += 13
		} else if (arr[i] > 'M' && arr[i] <= 'Z') || (arr[i] > 'm' && arr[i] <= 'z') {
			arr[i] -= 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
