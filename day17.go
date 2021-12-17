package main

import (
	"fmt"
)

type target struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

func main() {
	//target area: x=185..221, y=-122..-74
	t := target{185, 221, -122, -74}
	//t = target{20, 30, -10, -5}

	// Part1
	res := 0

	for x := 0; x < t.xmax; x++ {
		if (x*x+1)/2 < t.xmin {
			continue
		}
		for y := 0; y < 5000; y++ {
			hit, my := hitsTarget(x, y, t)
			if hit {
				if my > res {
					res = my
				}
			}
		}
	}

	fmt.Println(res)

	// Part2
	res = 0
	type c struct {
		x int
		y int
	}
	vels := make(map[c]bool)
	for x := 0; x <= t.xmax; x++ {
		for y := -5000; y < 5000; y++ {
			hit, _ := hitsTarget(x, y, t)
			if hit {
				vels[c{x, y}] = true
			}
		}
	}
	fmt.Println(len(vels))
}

func hitsTarget(dx int, dy int, t target) (bool, int) {
	x, y := 0, 0
	maxY := 0
	for true {
		x = x + dx
		y = y + dy
		if y > maxY {
			maxY = y
		}
		if x >= t.xmin && x <= t.xmax && y >= t.ymin && y <= t.ymax {
			return true, maxY
		}
		if dx > 0 {
			dx--
		}
		dy--
		if x > t.xmax || y < t.ymin {
			break
		}
	}

	return false, maxY
}
