package main

import (
	"bufio"
	"fmt"
	"os"

	//"strconv"
	"math"
	"sort"
	"strings"
)

func main() {
	f, _ := os.Open("day8_input")
	//f, _ := os.Open("day8_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	patterns := make([][]string, 0)
	digits := make([][]string, 0)
	for s.Scan() {
		spl := strings.Split(s.Text(), " | ")
		patterns = append(patterns, strings.Split(spl[0], " "))
		digits = append(digits, strings.Split(spl[1], " "))
	}

	// Part1
	res := 0
	for _, vs := range digits {
		for _, v := range vs {
			l := len(v)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				res++
			}
		}
	}

	fmt.Println(res)

	// Part2
	res = 0
	for ip, pattern := range patterns {
		digs := make(map[int][]string)
		for _, v := range pattern {
			digs[len(v)] = append(digs[len(v)], v)
		}
		assign := make(map[int]string)
		assign[1] = digs[2][0]
		assign[7] = digs[3][0]
		assign[4] = digs[4][0]
		assign[8] = digs[7][0]
		for _, d := range digs[6] {
			one := assign[1]
			four := assign[4]
			isNine := true
			for _, c := range four {
				if !strings.Contains(d, string(c)) {
					isNine = false
				}
			}
			if isNine {
				assign[9] = d
				continue
			}
			isZero := true
			for _, c := range one {
				if !strings.Contains(d, string(c)) {
					isZero = false
				}
			}
			if isZero {
				assign[0] = d
				continue
			}
			assign[6] = d
		}
		for _, d := range digs[5] {
			one := assign[1]
			isThree := true
			for _, c := range one {
				if !strings.Contains(d, string(c)) {
					isThree = false
				}
			}
			if isThree {
				assign[3] = d
				continue
			}
			six := assign[6]
			sim := 0
			for _, c := range d {
				if strings.Contains(six, string(c)) {
					sim++
				}
			}
			if sim == 5 {
				assign[5] = d
			} else {
				assign[2] = d
			}
		}
		rev := make(map[string]int)
		for k, v := range assign {
			rev[SortString(v)] = k
		}
		r := 0
		for i, d := range digits[ip] {
			n := rev[SortString(d)]
			r = r + n*int(math.Pow10(3-i))
		}
		res = res + r
	}
	fmt.Println(res)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
