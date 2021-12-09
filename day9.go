package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

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
			b := discBaisins(hmap, i, j)
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
}

func discBaisins(hmap [][]int, row int, col int) int {
	b := 0
	if hmap[row][col] == 9 {
		return b
	}
	b++
	hmap[row][col] = 9
	for n := -1; n < 2; n = n + 2 {
		if row+n >= 0 && row+n < len(hmap) {
			b = b + discBaisins(hmap, row+n, col)
		}
		if col+n >= 0 && col+n < len(hmap[row]) {
			b = b + discBaisins(hmap, row, col+n)
		}
	}

	return b
}
