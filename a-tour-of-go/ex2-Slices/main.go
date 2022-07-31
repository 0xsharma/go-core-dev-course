package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pix := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		pix[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			pix[i][j] = uint8(i ^ j + (i+j)/2)
		}
	}

	return pix
}

func main() {
	pic.Show(Pic)
}
