package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	m.Set(0, 0, color.RGBA{1, 253, 254, 255})
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
