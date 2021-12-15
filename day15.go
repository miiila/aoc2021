package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
	//"strconv"
	//"math"
	//"sort"
)

func main() {
	f, _ := os.Open("day15_input")
	//f, _ = os.Open("day15_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([][]int, 0)
	r := 0
	for s.Scan() {
		inputs = append(inputs, make([]int, 0))
		for _, v := range s.Text() {
			i, _ := strconv.ParseInt(string(v), 10, 0)
			inputs[r] = append(inputs[r], int(i))
		}
		r++
	}

	// Part1
	res := solve(inputs)
	fmt.Println(res)

	// Part2
	inputs2 := make([][]int, len(inputs)*5)
	for k := 0; k < 5; k++ {
		for i, v := range inputs {
			inputs2[i+len(inputs)*k] = make([]int, len(inputs[0])*5)
			for l := 0; l < 5; l++ {
				for j, w := range v {
					newV := w + l + k
					if newV > 9 {
						newV = newV%10 + 1
					}
					inputs2[i+len(inputs)*k][j+len(inputs[0])*l] = newV
				}
			}
		}
	}
	res = solve2(inputs2)
	fmt.Println(res)
}

func solve(inputs [][]int) int {
	resG := make([][]int, len(inputs))
	for i := 0; i < len(inputs); i++ {
		resG[i] = make([]int, len(inputs[0]))
	}

	for i, v := range inputs {
		for j, w := range v {
			if i < len(inputs) && j < len(inputs[i]) {
				if j < len(inputs)-1 {
					if resG[i][j+1] == 0 {
						resG[i][j+1] = resG[i][j] + w
					} else {
						if resG[i][j+1] > resG[i][j]+w {
							resG[i][j+1] = resG[i][j] + w
						}
					}
				}
				if i < len(inputs)-1 {
					if resG[i+1][j] == 0 {
						resG[i+1][j] = resG[i][j] + w
					} else {
						if resG[i+1][j] > resG[i][j]+w {
							resG[i+1][j] = resG[i][j] + w
						}
					}
				}
			}
		}
	}
	return resG[len(inputs)-1][len(inputs[0])-1] - inputs[0][0] + inputs[len(inputs)-1][len(inputs[0])-1]
}

func solve2(inputs [][]int) int {
	resG := make([][]int, len(inputs))
	for i := 0; i < len(inputs); i++ {
		resG[i] = make([]int, len(inputs[0]))
	}

	type coord struct {
		x int
		y int
	}

	stack := make([]coord, 0)
	stack = append(stack, coord{x: 0, y: 0})
	stackSet := make(map[coord]bool)
	//resG[0][0] = 9

	for len(stack) > 0 {
		c := stack[0]
		// MUST BE FIFO!!!!!
		stack = stack[1:]
		stackSet[c] = false
		i := c.x
		j := c.y
		w := inputs[i][j]
		if j < len(inputs)-1 {
			if resG[i][j+1] == 0 {
				resG[i][j+1] = resG[i][j] + w
				if !stackSet[coord{x: i, y: j + 1}] {
					stack = append(stack, coord{x: i, y: j + 1})
					stackSet[coord{x: i, y: j + 1}] = true
				}
			} else {
				if resG[i][j+1] > resG[i][j]+w {
					resG[i][j+1] = resG[i][j] + w
					if !stackSet[coord{x: i, y: j + 1}] {
						stack = append(stack, coord{x: i, y: j + 1})
						stackSet[coord{x: i, y: j + 1}] = true
					}
				}
			}
		}
		if j > 0 {
			if resG[i][j-1] == 0 {
				resG[i][j-1] = resG[i][j] + w
				if !stackSet[coord{x: i, y: j - 1}] {
					stack = append(stack, coord{x: i, y: j - 1})
					stackSet[coord{x: i, y: j - 1}] = true
				}
			} else {
				if resG[i][j-1] > resG[i][j]+w {
					resG[i][j-1] = resG[i][j] + w
					if !stackSet[coord{x: i, y: j - 1}] {
						stack = append(stack, coord{x: i, y: j - 1})
						stackSet[coord{x: i, y: j - 1}] = true
					}
				}
			}
		}
		if i < len(inputs)-1 {
			if resG[i+1][j] == 0 {
				resG[i+1][j] = resG[i][j] + w
				if !stackSet[coord{x: i + 1, y: j}] {
					stack = append(stack, coord{x: i + 1, y: j})
					stackSet[coord{x: i + 1, y: j}] = true
				}
			} else {
				if resG[i+1][j] > resG[i][j]+w {
					resG[i+1][j] = resG[i][j] + w
					if !stackSet[coord{x: i + 1, y: j}] {
						stack = append(stack, coord{x: i + 1, y: j})
						stackSet[coord{x: i + 1, y: j}] = true
					}
				}
			}
		}
		if i > 0 {
			if resG[i-1][j] == 0 {
				resG[i-1][j] = resG[i][j] + w
				if !stackSet[coord{x: i - 1, y: j}] {
					stack = append(stack, coord{x: i - 1, y: j})
					stackSet[coord{x: i - 1, y: j}] = true
				}
			} else {
				if resG[i-1][j] > resG[i][j]+w {
					resG[i-1][j] = resG[i][j] + w
					if !stackSet[coord{x: i - 1, y: j}] {
						stack = append(stack, coord{x: i - 1, y: j})
						stackSet[coord{x: i - 1, y: j}] = true
					}
				}
			}
		}
	}
	return resG[len(inputs)-1][len(inputs[0])-1] - inputs[0][0] + inputs[len(inputs)-1][len(inputs[0])-1]
}
