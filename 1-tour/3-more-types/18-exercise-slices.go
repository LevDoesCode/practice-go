package main

import "pic"

func Pic(dx, dy int) [][]uint8 {
	myPic := make([]([]uint8), dy)
	for i := range myPic {
		myPic[i] = make([]uint8, dx)
		for j := range myPic[i] {
			myPic[i][j] = uint8((i + j) / 2)
		}
	}

	return myPic
}

func main() {
	pic.Show(Pic)
}
