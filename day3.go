package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

func main() {
	f, _ := os.Open("day3_input")
	s := bufio.NewScanner(f)
	inputs := make([]string, 0)
	for s.Scan() {
		inputs = append(inputs, s.Text())
	}

	// Part1
	g := make([]int, len(inputs[0]))
	for _, v := range inputs {
		for i := 0; i < len(v); i++ {
			if v[i] == '0' {
				g[i]--
			} else {
				g[i]++
			}
		}
	}

	var gs, es string
	for _, v := range g {
		if v < 0 {
			gs += "0"
			es += "1"
		} else {
			gs += "1"
			es += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gs, 2, 64)
	epsilon, _ := strconv.ParseInt(es, 2, 64)
	fmt.Println(gamma * epsilon)

	// Part2
	pos := 0
	ox := inputs
	for len(ox) > 1 {
		cb := mostCommonBit(ox, pos, '1')
		filt := make([]string, 0)
		for _, v := range ox {
			if v[pos] == cb {
				filt = append(filt, v)
			}
		}
		ox = filt
		pos++
	}

	co := inputs
	pos = 0
	for len(co) > 1 {
		var cb byte
		if mostCommonBit(co, pos, '1') == '0' {
			cb = '1'
		} else {
			cb = '0'
		}
		filt := make([]string, 0)
		for _, v := range co {
			if v[pos] == cb {
				filt = append(filt, v)
			}
		}
		co = filt
		pos++
	}
	oxv, _ := strconv.ParseInt(ox[0], 2, 64)
	cov, _ := strconv.ParseInt(co[0], 2, 64)
	fmt.Println(oxv * cov)

}

func mostCommonBit(inputs []string, pos int, def byte) byte {
	b := 0
	for _, v := range inputs {
		if v[pos] == '0' {
			b--
		} else {
			b++
		}
	}

	if b > 0 {
		return '1'
	}
	if b < 0 {
		return '0'
	}

	return def
}
