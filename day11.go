package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	//"math"
	//"sort"
	//"strings"
)

func main() {
	f, _ := os.Open("day11_input")
	//f, _ = os.Open("day11_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	var grid [10][10]int

	gi := 0
	for s.Scan() {
		var gr [10]int
		for i, v := range s.Text() {
			n, _ := strconv.ParseInt(string(v), 10, 0)
			gr[i] = int(n)
		}
		grid[gi] = gr
		gi++
	}

	// Part1
	res := 0

	for i := 0; i < 100; i++ {
		flashed := make(map[[2]int]bool)
		toFlash := make([][2]int, 0)
		// raise energy
		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				grid[r][c]++
				if grid[r][c] > 9 {
					toFlash = append(toFlash, [2]int{r, c})
				}
			}
		}
		lf := len(toFlash)
		for lf > 0 {
			c := toFlash[0]
			toFlash = toFlash[1:]
			if flashed[c] {
				lf--
				continue
			}
			res++
			grid[c[0]][c[1]] = 0
			flashed[c] = true
			// adjacent
			for rd := -1; rd < 2; rd++ {
				for cd := -1; cd < 2; cd++ {
					row := c[0] + rd
					col := c[1] + cd
					if (rd == 0 && cd == 0) || (flashed[[2]int{row, col}]) || row < 0 || row > 9 || col < 0 || col > 9 {
						continue
					}
					grid[row][col]++
					if grid[row][col] > 9 {
						toFlash = append(toFlash, [2]int{row, col})
					}
				}
			}
			lf = len(toFlash)
		}

	}
	fmt.Println(res)

	// Part2
	res = 100
outer:
	for true {
		flashed := make(map[[2]int]bool)
		toFlash := make([][2]int, 0)
		done := 0
		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				done = done + grid[r][c]
			}
		}
		if done == 0 {
			break outer
		}
		// raise energy
		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				grid[r][c]++
				if grid[r][c] > 9 {
					toFlash = append(toFlash, [2]int{r, c})
				}
			}
		}
		lf := len(toFlash)
		for lf > 0 {
			c := toFlash[0]
			toFlash = toFlash[1:]
			if flashed[c] {
				lf = len(toFlash)
				continue
			}
			grid[c[0]][c[1]] = 0
			flashed[c] = true
			// adjacent
			for rd := -1; rd < 2; rd++ {
				for cd := -1; cd < 2; cd++ {
					row := c[0] + rd
					col := c[1] + cd
					if (rd == 0 && cd == 0) || (flashed[[2]int{row, col}]) || row < 0 || row > 9 || col < 0 || col > 9 {
						continue
					}
					grid[row][col]++
					if grid[row][col] > 9 {
						toFlash = append(toFlash, [2]int{row, col})
					}
				}
			}
			lf = len(toFlash)
		}
		res++
	}
	fmt.Println(res)
}
