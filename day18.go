package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
	//"math"
	//"sort"
)

type pair struct {
	firstPair  *pair
	secondPair *pair
	firstVal   int
	secondVal  int
}

func main() {
	f, _ := os.Open("day18_input")
	//f, _ = os.Open("day18_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([]string, 0)
	for s.Scan() {
		inputs = append(inputs, s.Text())
	}

	// Part1
	res := 0

	var resPair *pair
	tmp := parseInput(inputs[0])
	resPair = &tmp
	for _, v := range inputs[1:] {
		tmp := parseInput(v)
		np := add(resPair, &tmp)
		resPair = np
	}
	res = countMagnitude(resPair)
	fmt.Println(res)

	// Part2
	res = 0
	for i, v := range inputs {
		for j := i + 1; j < len(inputs); j++ {
			first1, first2 := parseInput(v), parseInput(v)
			second1, second2 := parseInput(inputs[j]), parseInput(inputs[j])
			mag1 := countMagnitude(add(&first1, &second1))
			mag2 := countMagnitude(add(&second2, &first2))
			if mag1 > res {
				res = mag1
			}
			if mag2 > res {
				res = mag2
			}
		}
	}
	fmt.Println(res)
}

func countMagnitude(p *pair) int {
	mag := 0
	if p.firstVal > -1 {
		mag = mag + 3*p.firstVal
	}
	if p.secondVal > -1 {
		mag = mag + 2*p.secondVal
	}
	if p.firstPair != nil {
		mag = mag + 3*countMagnitude(p.firstPair)
	}
	if p.secondPair != nil {
		mag = mag + 2*countMagnitude(p.secondPair)
	}

	return mag
}

func add(p1 *pair, p2 *pair) *pair {
	np := pair{firstPair: p1, secondPair: p2, firstVal: -1, secondVal: -1}
	reduce(&np)
	return &np
}

func reduce(p *pair) {
	run := true
	for run {
		run = explode(p, 1)
		if run {
			continue
		}
		run = split(p)
	}
}

func split(p *pair) bool {
	if p.firstVal > 9 {
		fv := p.firstVal / 2
		np := pair{firstVal: fv, secondVal: p.firstVal - fv}
		p.firstVal = -1
		p.firstPair = &np
		return true
	}
	didSplit := false
	if p.firstPair != nil {
		didSplit = split(p.firstPair)
	}
	if didSplit {
		return true
	}
	if p.secondVal > 9 {
		fv := p.secondVal / 2
		np := pair{firstVal: fv, secondVal: p.secondVal - fv}
		p.secondVal = -1
		p.secondPair = &np
		return true
	}
	if p.secondPair != nil {
		didSplit = split(p.secondPair)
	}

	return didSplit
}

func explode(p *pair, level int) bool {
	l, r := -1, -1
	pt := make([]*pair, 0)
	pt = append(pt, p)
	ex, found := findExplode(p, 0, pt)
	if found {
		lDone, rDone := false, false
		l, r = ex[4].firstVal, ex[4].secondVal
		if ex[3].firstPair == ex[4] {
			ex[3].firstPair = nil
			ex[3].firstVal = 0
			if ex[3].secondPair != nil {
				addToRight(ex[3].secondPair, r)
			} else {
				ex[3].secondVal = ex[3].secondVal + r
			}
			rDone = true
		}
		if ex[3].secondPair == ex[4] {
			ex[3].secondPair = nil
			ex[3].secondVal = 0
			if ex[3].firstPair != nil {
				addToLeft(ex[3].firstPair, l)
			} else {
				ex[3].firstVal = ex[3].firstVal + l
			}
			lDone = true
		}
		for i := 3; i >= 1; i-- {
			if ex[i-1].firstPair == ex[i] && !rDone {
				if ex[i-1].secondPair != nil {
					addToRight(ex[i-1].secondPair, r)
				} else {
					ex[i-1].secondVal = ex[i-1].secondVal + r
				}
				rDone = true
			}
			if ex[i-1].secondPair == ex[i] && !lDone {
				if ex[i-1].firstPair != nil {
					addToLeft(ex[i-1].firstPair, l)
				} else {
					ex[i-1].firstVal = ex[i-1].firstVal + l
				}
				lDone = true
			}
		}
	}
	return found
}

func findExplode(p *pair, level int, path []*pair) ([]*pair, bool) {
	if level == 4 {
		return path, true
	}
	if p.firstPair != nil {
		newPath := append(path, p.firstPair)
		pt, fp := findExplode(p.firstPair, level+1, newPath)
		if fp {
			return pt, true
		}
	}
	if p.secondPair != nil {
		newPath := append(path, p.secondPair)
		pt, sp := findExplode(p.secondPair, level+1, newPath)
		if sp {
			return pt, true
		}
	}

	return path, false
}

func addToLeft(p *pair, val int) {
	if p.secondPair != nil {
		addToLeft(p.secondPair, val)
	} else {
		p.secondVal = p.secondVal + val
	}
}

func addToRight(p *pair, val int) {
	if p.firstPair != nil {
		addToRight(p.firstPair, val)
	} else {
		p.firstVal = p.firstVal + val
	}
}

func parseInput(input string) pair {
	stack := make([]*pair, 0)
	isFirst := true
	var outerPair pair
	for len(input) > 0 {
		c := input[0]
		input = input[1:]
		switch c {
		case '[':
			p := pair{firstVal: -1, secondVal: -1}
			if len(stack) > 0 {
				if isFirst {
					stack[len(stack)-1].firstPair = &p
				} else {
					stack[len(stack)-1].secondPair = &p
				}

			}
			isFirst = true
			stack = append(stack, &p)
		case ']':
			outerPair = *stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			isFirst = true
		case ',':
			isFirst = false
		default:
			ci, _ := strconv.ParseInt(string(c), 10, 0)
			if isFirst {
				stack[len(stack)-1].firstVal = int(ci)
			} else {
				stack[len(stack)-1].secondVal = int(ci)
			}
		}

	}

	return outerPair
}

func printPair(p *pair, out string) string {

	if p.firstVal > -1 {
		out = out + strconv.FormatInt(int64(p.firstVal), 10) + ","
	}
	if p.firstPair != nil {
		out = out + "["
		out = out + printPair(p.firstPair, "")
		out = out + ","

	}

	if p.secondVal > -1 {
		out = out + strconv.FormatInt(int64(p.secondVal), 10) + "]"
	}
	if p.secondPair != nil {
		out = out + "["
		out = out + printPair(p.secondPair, "")
		out = out + "]"
	}

	return out
}
