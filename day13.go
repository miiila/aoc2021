package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	//"math"
	//"sort"
	"strings"
)

type fold struct {
	axes rune
	val  int
}

func main() {
	f, _ := os.Open("day13_input")
	//f, _ = os.Open("day13_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	xs := make(map[int]map[int]bool, 0)
	ys := make(map[int]map[int]bool, 0)
	parseIns := false
	instructions := make([]fold, 0)
	for s.Scan() {
		if s.Text() == "" {
			parseIns = true
			continue
		}
		if !parseIns {
			xy := strings.Split(s.Text(), ",")
			x, _ := strconv.ParseInt(xy[0], 10, 0)
			y, _ := strconv.ParseInt(xy[1], 10, 0)
			_, ok := xs[int(x)]
			if !ok {
				xs[int(x)] = make(map[int]bool)
			}
			xs[int(x)][int(y)] = true
			_, ok = ys[int(y)]
			if !ok {
				ys[int(y)] = make(map[int]bool)
			}
			ys[int(y)][int(x)] = true
		} else {
			i := strings.Split(s.Text(), "=")
			v, _ := strconv.ParseInt(i[1], 10, 0)
			ins := fold{axes: rune(i[0][len(i[0])-1]), val: int(v)}
			instructions = append(instructions, ins)
		}
	}

	// Part1
	res := 0

	ins := instructions[0]
	if ins.axes == 'x' {
		for _, v := range foldBy(xs, ins.val) {
			res += len(v)
		}
	} else {
		for _, v := range foldBy(ys, ins.val) {
			res += len(v)
		}
	}
	fmt.Println(res)

	// Part2
	res = 0
	for _, ins := range instructions {
		if ins.axes == 'x' {
			newXs := foldBy(xs, ins.val)
			newYs := make(map[int]map[int]bool)
			for k, v := range newXs {
				for w, _ := range v {
					if _, ok := newYs[w]; !ok {
						newYs[w] = make(map[int]bool)
					}
					newYs[w][k] = true
				}
			}
			xs, ys = newXs, newYs
		} else {
			newYs := foldBy(ys, ins.val)
			newXs := make(map[int]map[int]bool)
			for k, v := range newYs {
				for w, _ := range v {
					if _, ok := newXs[w]; !ok {
						newXs[w] = make(map[int]bool)
					}
					newXs[w][k] = true
				}
			}
			xs, ys = newXs, newYs
		}
	}

	var resG [100][100]rune
	for k, v := range xs {
		for w, _ := range v {
			resG[k][w] = '*'
		}
	}

	for i := 0; i < 7; i++ {
		for j := 0; j < 50; j++ {
			if resG[j][i] == '*' {
				print("*")
			} else {
				print(" ")
			}
		}
		print("\n")
	}

}

func foldBy(xs map[int]map[int]bool, along int) map[int]map[int]bool {
	newXS := make(map[int]map[int]bool)
	for k, v := range xs {
		if k <= along {
			newXS[k] = v
			continue
		}
	}
	for k, v := range xs {
		if k <= along {
			continue
		}
		newX := along + along - k
		for w, _ := range v {
			_, ok := newXS[newX]
			if !ok {
				newXS[newX] = make(map[int]bool)
			}
			newXS[newX][w] = true
		}
	}
	return newXS
}
