package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"strings"
	"regexp"
	"strconv"
	//"math"
	//"sort"
  "github.com/ungerik/go3d/vec3"
)

type inst struct {
	onoff bool
	xMin  int
	xMax  int
	yMin  int
	yMax  int
	zMin  int
	zMax  int
}

type cubeCoord struct {
	x int
	y int
	z int
}

type cuboid struct {
  xMin int
	xMax  int
	yMin  int
	yMax  int
	zMin  int
	zMax  int
}

func main() {
	f, _ := os.Open("day22_input")
	//f, _ = os.Open("day_input_test")
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
		ins := inst{xMin: int(xMin), xMax: int(xMax), yMin: int(yMin), yMax: int(yMax), zMin: int(zMin), zMax: int(zMax)}
		if r[1] == "on" {
			ins.onoff = true
		} else {
			ins.onoff = false
		}
		inputs = append(inputs, ins)
	}

	// Part1
	res := 0
	state := make(map[cubeCoord] bool)
	//for _, ins := range inputs {
		//for x := ins.xMin; x <= ins.xMax; x++ {
			//for y := ins.yMin; y <= ins.yMax; y++ {
				//for z := ins.zMin; z <= ins.zMax; z++ {
					//state[cubeCoord{x, y, z}] = ins.onoff
				//}
			//}
		//}
	//}
  for _, v := range state {
    if v {
      res++
    }
  }
	fmt.Println(res)

	// Part2
	res = 0
  c1 := cuboid{10,12,10,12,10,12}
  c2 := cuboid{11,13,11,13,11,13}
	fmt.Println(join(c1,c2))
	fmt.Println(res)
}

func join(c1 cuboid, c2 cuboid) vec3.Box {
    t11 := vec3.T{float32(c1.xMin), float32(c1.yMin), float32(c1.zMin)}
    t12 := vec3.T{float32(c1.xMax), float32(c1.yMax), float32(c1.zMax)}
    b1 := vec3.Box{t11, t12}
    t21 := vec3.T{float32(c2.xMin), float32(c2.yMin), float32(c2.zMin)}
    t22 := vec3.T{float32(c2.xMax), float32(c2.yMax), float32(c2.zMax)}
    b2 := vec3.Box{t21, t22}

    return vec3.Joined(&b1,&b2)
}
