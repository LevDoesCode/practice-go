package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	// one array for x value, one array for y values, array of integer for pixel RGBA values
	image          [][][]uint8
	x0, y0, x1, y1 int
}

func (img *Image) NewImage(x, y int) *Image {
	image := make([][][]uint8, y)
	for i := 0; i < x; i++ {
		image[i] = make([][]uint8, x)
		for j := 0; j < y; j++ {
			image[i][j] = make([]uint8, 4)
		}
	}
	return &Image{image, 0, 0, x, y}
}

func (p *Image) ColorModel() color.Model { return color.RGBAModel }

func (p *Image) At(x, y int) color.Color {
	return color.RGBA{
		p.image[x][y][0],
		p.image[x][y][1],
		p.image[x][y][2],
		p.image[x][y][3],
	}
}

func (p *Image) Bounds() image.Rectangle {
	return image.Rect(p.x0, p.y0, p.x1, p.y1)
}

func (p *Image) crazyColors() *Image {
	for i := 0; i < len(p.image)-1; i++ {
		for j := 0; j < len(p.image[i])-1; j++ {
			p.image[i][j][0] = uint8(i + j/2)
			p.image[i][j][1] = uint8(i + j*2)
			p.image[i][j][2] = uint8(i + j*2)
			p.image[i][j][3] = uint8(255)
		}
	}
	return p
}

func main() {
	var img *Image
	newImg := img.NewImage(255, 255)
	newImg.crazyColors()
	pic.ShowImage(newImg)
}

/////////////////////////////
// Another solution

// package main

// import (
//     "golang.org/x/tour/pic"
//     "image/color"
//     "image"
// )

// type Image struct{
//     Rect image.Rectangle
// }

// func (image Image) ColorModel() color.Model {
//         return color.RGBAModel
// }
// func (i Image) Bounds() image.Rectangle {
//     return i.Rect
// }
// func (image Image) At(x, y int) color.Color {
//     return color.RGBA{uint8(x*y), uint8(y+x), 255, 255}
// }

// func main() {
//     m := Image{image.Rect(0,0,300,300)}
//     pic.ShowImage(m)
// }
//////////////////////////////////////
// Another solution
// package main

// import (
//     "golang.org/x/tour/pic"
//     "image"
// )

// type Image struct{
//     image.RGBA
// }

// func (i Image) crazyColors() Image {
//     for p := 0; p < len(i.Pix); p += 4 {
//         i.Pix[p] = uint8(p + p*3)
//         i.Pix[p+1] = uint8(p + p*4)
//         i.Pix[p+2] = uint8(p + p*5)
//         i.Pix[p+3] = uint8(p + p*2)
//     }
//     return i
// }

// func main() {
//     rect := image.Rect(0,0,255,255)
//     m := Image{*image.NewRGBA(rect)}
//     m.crazyColors()
//     pic.ShowImage(&m)
// }
