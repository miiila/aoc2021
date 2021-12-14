package main

import (
	"bufio"
	"fmt"
	"os"

	//"strconv"
	//"math"
	//"sort"
	"strings"
)

func main() {
	f, _ := os.Open("day14_input")
	//f, _ = os.Open("day14_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	rules := make(map[string]string)
	s.Scan()
	input := s.Text()
	s.Scan()
	for s.Scan() {
		rule := strings.Split(s.Text(), " -> ")
		rules[rule[0]] = rule[1]
	}

	// Part1
	origInput := input
	res := 0
	for i := 0; i < 10; i++ {
		newInput := string(input[0])
		for j := 0; j < len(input)-1; j++ {
			newInput = newInput + rules[string(input[j])+string(input[j+1])] + string(input[j+1])
		}
		input = newInput
	}

	counts := make(map[rune]int)

	for _, r := range input {
		counts[r]++
	}

	minCount, maxCount := 1000000, 0

	for _, v := range counts {
		if v < minCount {
			minCount = v
		}
		if v > maxCount {
			maxCount = v
		}
	}
	res = maxCount - minCount
	fmt.Println(res)

	// Part2

	res = 0
	input = origInput
	counts2 := make(map[string]int)
	for _, r := range input {
		counts2[string(r)]++
	}
	m := make(map[string]int)
	for j := 0; j < len(input)-1; j++ {
		m[string(input[j])+string(input[j+1])]++
	}
	for i := 0; i < 40; i++ {
		nm := make(map[string]int)
		for k, v := range m {
			counts2[rules[k]] = v + counts2[rules[k]]
			a := string(k[0]) + rules[k]
			b := rules[k] + string(k[1])
			nm[string(a)] = nm[string(a)] + v
			nm[string(b)] = nm[string(b)] + v
		}
		m = nm
	}

	minCount, maxCount = 1000000000000000000, 0

	for _, v := range counts2 {
		if v < minCount {
			minCount = v
		}
		if v > maxCount {
			maxCount = v
		}
	}
	res = maxCount - minCount
	fmt.Println(res)
}
