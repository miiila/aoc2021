package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	//"strconv"
	//"math"
	//"sort"
)

type coord struct {
	x int
	y int
}

func main() {
	f, _ := os.Open("day20_input")
  f, _ = os.Open("day20_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	//inputs := make([]int, 0)
	s.Scan()
	alg := s.Text()
	s.Scan()
	image := make(map[coord]int)
	var max coord
	x := 0
	for s.Scan() {
		for y, r := range s.Text() {
			if r == '#' {
				image[coord{x, y}] = 1
			}
			max = coord{x, y}
		}
		x++
	}
	min := coord{0, 0}

	// Part1
	res := 0

	for i := 0; i < 50; i++ {
		ni, nmin, nmax := iterate(alg, image, min, max, i%2)
		image, min, max = ni, nmin, nmax
		//printImage(image, min, max)
	}
		fmt.Println(len(image))

	// Part2
	res = 0
	fmt.Println(res)
}

func iterate(alg string, image map[coord]int, min coord, max coord, iter int) (map[coord]int, coord, coord) {
	bbTL := coord{min.x - 1, min.y - 1}
	bbBR := coord{max.x + 1, max.y + 1}

	newImage := make(map[coord]int)
	for x := bbTL.x; x <= bbBR.x; x++ {
		for y := bbTL.y; y <= bbBR.y; y++ {
			v := 0
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					v = v << 1
          if x+i < bbTL.x+1 || y+j < bbTL.y+1 || x+i > bbBR.x-1 || y+j > bbBR.y-1 {
						v = v | iter
					} else {
						//fmt.Println(x+i, y+j, v, image[coord{x + i, y + j}])
						if v&1 != 0 {
							panic("Boom")
						}
            v = v | image[coord{x + i, y + j}]
					}
				}
			}
			//fmt.Println(x, y, v)
			if alg[v] == '#' {
				newImage[coord{x, y}] = 1
			}

		}
	}

	return newImage, bbTL, bbBR

}

func printImage(image map[coord]int, min coord, max coord) {
	res := ""
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			if image[coord{x, y}] == 1 {
				res = res + "#"
			} else {
				res = res + "."
			}
		}
		res = res + "\n"
	}
	fmt.Println(res)
}
