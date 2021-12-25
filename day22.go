package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"strings"
	"math"
	"regexp"
	"strconv"
	//"sort"
	"github.com/ungerik/go3d/vec3"
)

type inst struct {
	onoff  bool
	cuboid cuboid
}

type cubeCoord struct {
	x int
	y int
	z int
}

type cuboid struct {
	xMin int
	xMax int
	yMin int
	yMax int
	zMin int
	zMax int
}

func (c cuboid) getSize() int {
	return (c.xMax - c.xMin + 1) * (c.yMax - c.yMin + 1) * (c.zMax - c.zMin + 1)
}

func (c1 *cuboid) intersects(c2 *cuboid) bool {
	t11 := vec3.T{float32(c1.xMin), float32(c1.yMin), float32(c1.zMin)}
	t12 := vec3.T{float32(c1.xMax), float32(c1.yMax), float32(c1.zMax)}
	b1 := vec3.Box{t11, t12}
	t21 := vec3.T{float32(c2.xMin), float32(c2.yMin), float32(c2.zMin)}
	t22 := vec3.T{float32(c2.xMax), float32(c2.yMax), float32(c2.zMax)}
	b2 := vec3.Box{t21, t22}

	return b1.Intersects(&b2)
}

func main() {
	f, _ := os.Open("day22_input")
	//f, _ = os.Open("day22_input_test")
	uBound := 20
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([]inst, 0)
	for s.Scan() {
		re := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
		r := re.FindStringSubmatch(s.Text())
		xMin, _ := strconv.ParseInt(r[2], 10, 0)
		xMax, _ := strconv.ParseInt(r[3], 10, 0)
		yMin, _ := strconv.ParseInt(r[4], 10, 0)
		yMax, _ := strconv.ParseInt(r[5], 10, 0)
		zMin, _ := strconv.ParseInt(r[6], 10, 0)
		zMax, _ := strconv.ParseInt(r[7], 10, 0)
		ins := inst{cuboid: cuboid{xMin: int(xMin), xMax: int(xMax), yMin: int(yMin), yMax: int(yMax), zMin: int(zMin), zMax: int(zMax)}}
		if r[1] == "on" {
			ins.onoff = true
		} else {
			ins.onoff = false
		}
		inputs = append(inputs, ins)
	}

	// Part1
	res := 0
	state := make(map[cubeCoord]bool)
	for _, ins := range inputs[0:uBound] {
		for x := ins.cuboid.xMin; x <= ins.cuboid.xMax; x++ {
			for y := ins.cuboid.yMin; y <= ins.cuboid.yMax; y++ {
				for z := ins.cuboid.zMin; z <= ins.cuboid.zMax; z++ {
					state[cubeCoord{x, y, z}] = ins.onoff
				}
			}
		}
	}
	for _, v := range state {
		if v {
			res++
		}
	}
	fmt.Println(res)

	// Part2
	res = 0
	results := make([]inst, 0)
	for _, ins := range inputs {
		for _, c := range results {
			if ins.cuboid.intersects(&c.cuboid) {
				modC := inst{}
				modC.cuboid = getIntersection(ins.cuboid, c.cuboid)
				if ins.onoff && c.onoff {
					modC.onoff = false
				} else if !ins.onoff && !c.onoff {
					modC.onoff = true
				} else {
					modC.onoff = ins.onoff
				}
				results = append(results, modC)
			}
		}
		if ins.onoff {
			results = append(results, ins)
		}
	}

	for _, i := range results {
		if i.onoff {
			res = res + i.cuboid.getSize()
		} else {
			res = res - i.cuboid.getSize()
		}
	}
	fmt.Println(res)
}

func getIntersection(c1 cuboid, c2 cuboid) cuboid {
	t11 := vec3.T{float32(c1.xMin), float32(c1.yMin), float32(c1.zMin)}
	t12 := vec3.T{float32(c1.xMax), float32(c1.yMax), float32(c1.zMax)}
	b1 := vec3.Box{t11, t12}
	t21 := vec3.T{float32(c2.xMin), float32(c2.yMin), float32(c2.zMin)}
	t22 := vec3.T{float32(c2.xMax), float32(c2.yMax), float32(c2.zMax)}
	b2 := vec3.Box{t21, t22}

	if b1.Intersects(&b2) {
		return cuboid{xMin: intMax(c1.xMin, c2.xMin), yMin: intMax(c1.yMin, c2.yMin), zMin: intMax(c1.zMin, c2.zMin), xMax: intMin(c1.xMax, c2.xMax), yMax: intMin(c1.yMax, c2.yMax), zMax: intMin(c1.zMax, c2.zMax)}
	}
	panic("Boom")

}

func intMax(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func intMin(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
