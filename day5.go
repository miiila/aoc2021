package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int64
	y int64
}

func main() {
	f, _ := os.Open("day5_input")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([][]string, 0)
	for s.Scan() {
		inputs = append(inputs, strings.Split(s.Text(), " -> "))
	}

	res := 0

	// Part1
	plan := make(map[point]int64)
	for _, v := range inputs {
		from := strings.Split(v[0], ",")
		to := strings.Split(v[1], ",")
		if from[0] == to[0] || from[1] == to[1] {
			xf, _ := strconv.ParseInt(from[0], 10, 0)
			yf, _ := strconv.ParseInt(from[1], 10, 0)
			xt, _ := strconv.ParseInt(to[0], 10, 0)
			yt, _ := strconv.ParseInt(to[1], 10, 0)
			if xf == xt {
				for i := intMin(yf, yt); i <= intMax(yf, yt); i++ {
					p := point{x: xf, y: i}
					plan[p]++
				}
			} else {
				for i := intMin(xf, xt); i <= intMax(xf, xt); i++ {
					p := point{x: i, y: yf}
					plan[p]++
				}
			}
		}
	}
	for _, v := range plan {
		if v > 1 {
			res++
		}
	}
	fmt.Println(res)

	// Part2
	for _, v := range inputs {
		from := strings.Split(v[0], ",")
		to := strings.Split(v[1], ",")
		if from[0] != to[0] && from[1] != to[1] {
			xf, _ := strconv.ParseInt(from[0], 10, 0)
			yf, _ := strconv.ParseInt(from[1], 10, 0)
			xt, _ := strconv.ParseInt(to[0], 10, 0)
			yt, _ := strconv.ParseInt(to[1], 10, 0)
			x, y := xf, yf
			for x != xt && y != yt {
				p := point{x: x, y: y}
				plan[p]++
				if xf > xt {
					x--
				} else {
					x++
				}
				if yf > yt {
					y--
				} else {
					y++
				}
			}
			p := point{x: xt, y: yt}
			plan[p]++
		}
	}
	res = 0
	for _, v := range plan {
		if v > 1 {
			res++
		}
	}
	fmt.Println(res)
}

func intMax(a int64, b int64) int64 {
	return int64(math.Max(float64(a), float64(b)))
}

func intMin(a int64, b int64) int64 {
	return int64(math.Min(float64(a), float64(b)))
}
