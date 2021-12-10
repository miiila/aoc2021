package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"math"
	"sort"
	//"strings"
)

func main() {
	f, _ := os.Open("day10_input")
	//f, _ := os.Open("day10_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	chunks := make([]string, 0)
	for s.Scan() {
		chunks = append(chunks, s.Text())
	}
	points := make(map[rune]int)
	closing := make(map[rune]int)
	points[')'] = 3
	points[']'] = 57
	points['}'] = 1197
	points['>'] = 25137
	closing['('] = ')'
	closing['['] = ']'
	closing['{'] = '}'
	closing['<'] = '>'
	// Part1
	res := 0
	for _, line := range chunks {
		expected := make([]rune, 0)
		for _, r := range line {
			if r == '(' || r == '[' || r == '{' || r == '<' {
				expected = append(expected, rune(closing[r]))
				continue
			}
			l := len(expected) - 1
			if r == expected[l] {
				expected = expected[0:l]
				continue
			}

			res = res + points[r]
			break
		}
	}

	fmt.Println(res)

	// Part2
	res2 := make([]int, 0)
	points[')'] = 1
	points[']'] = 2
	points['}'] = 3
	points['>'] = 4

	for _, line := range chunks {
		expected := make([]rune, 0)
		correct := true
		for _, r := range line {
			if r == '(' || r == '[' || r == '{' || r == '<' {
				expected = append(expected, rune(closing[r]))
				continue
			}
			l := len(expected) - 1
			if r == expected[l] {
				expected = expected[0:l]
				continue
			}
			correct = false
			break
		}
		if correct {
			score := 0
			for i := len(expected) - 1; i >= 0; i-- {
				score = score*5 + points[expected[i]]
			}
			res2 = append(res2, score)
		}
	}

	sort.Sort(sort.IntSlice(res2))
	fmt.Println(res2[len(res2)/2])
}
