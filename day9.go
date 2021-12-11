package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"image"
	"image/color"
	"image/gif"
	"strconv"
	//"math"
	//"sort"
	//"strings"
)

func main() {
	f, _ := os.Open("day9_input")
	//f, _ := os.Open("day8_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	hmap := make([][]int, 0)
	for s.Scan() {
		t := s.Text()
		row := make([]int, 0)
		for _, v := range t {
			i, _ := strconv.ParseInt(string(v), 10, 0)
			row = append(row, int(i))
		}
		hmap = append(hmap, row)
	}

	// image
	width := len(hmap[0])
	height := len(hmap)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	pal := color.Palette{color.RGBA{0, 0, 0, 0xff}}
	for i := 8; i >= 0; i-- {
		c := color.RGBA{255, 255 - uint8(255/8*i), 0, 0xff}
		pal = append(pal, c)
	}
	img := image.NewPaletted(image.Rectangle{upLeft, lowRight}, pal)
	myGif := gif.GIF{Image: make([]*image.Paletted, 0)}

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	black := color.RGBA{0, 0, 0, 0xff}
	//white := color.RGBA{255, 255, 255, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if hmap[x][y] == 9 {
				img.Set(x, y, black)
			} else {
				c := color.RGBA{255, 255 - uint8(255/8*hmap[x][y]), 0, 0xff}
				img.Set(x, y, c)
			}
		}
	}

	// Encode as GIF.
	myGif.Image = append(myGif.Image, img)
	myGif.Delay = append(myGif.Delay, 10)

	// Part1
	res := 0
	for i, v := range hmap {
		for j, w := range v {
			if w == 9 {
				continue
			}
			s := true
			for n := -1; n < 2; n = n + 2 {
				if i+n >= 0 && i+n < len(hmap) {
					if hmap[i+n][j] < w {
						s = false
					}
				}
				if j+n >= 0 && j+n < len(v) {
					if hmap[i][j+n] < w {
						s = false
					}
				}
			}
			if s {
				res = res + 1 + w
			}
		}
	}
	fmt.Println(res)

	// Part2
	res = 0
	baisins := make([]int, 0)
	for i := 0; i < len(hmap); {
		for j := 0; j < len(hmap[i]); {
			if hmap[i][j] == 9 {
				j++
				continue
			}
			b := discBaisins(hmap, i, j, &myGif)
			baisins = append(baisins, b)
			j++
		}
		i++
	}
	sort.Slice(baisins, func(i, j int) bool {
		return baisins[i] > baisins[j]
	})
	res = baisins[0] * baisins[1] * baisins[2]
	fmt.Println(res)

	ff, _ := os.Create("basins.gif")
	//fmt.Println(myGif.Image)
	e := gif.EncodeAll(ff, &myGif)
	fmt.Println(e)
}

func discBaisins(hmap [][]int, row int, col int, myGif *gif.GIF) int {
	b := 0
	if hmap[row][col] == 9 {
		return b
	}
	b++
	hmap[row][col] = 9
	// Set color for each pixel.
	black := color.RGBA{0, 0, 0, 0xff}
	img := myGif.Image[len(myGif.Image)-1]
	var cImg image.Paletted
	cImg = image.Paletted{Pix: make([]uint8, len(img.Pix)), Stride: img.Stride, Rect: img.Rect, Palette: img.Palette}
	copy(cImg.Pix, img.Pix)
	cImg.Set(row, col, black)
	myGif.Image = append(myGif.Image, &cImg)
	myGif.Delay = append(myGif.Delay, 1/10000)
	for n := -1; n < 2; n = n + 2 {
		if row+n >= 0 && row+n < len(hmap) {
			b = b + discBaisins(hmap, row+n, col, myGif)
		}
		if col+n >= 0 && col+n < len(hmap[row]) {
			b = b + discBaisins(hmap, row, col+n, myGif)
		}
	}

	return b
}
