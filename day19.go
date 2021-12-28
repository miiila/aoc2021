package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"strconv"
	"math"
	//"sort"
)

type coord struct {
	x int
	y int
	z int
}

type scannerT []coord

func applySymmetry(which int, x int, y int, z int) (int, int, int) {
	var t int
	// Peel off the "do we swap the tetrahedrons?" bit
	if which&16 == 0 {
		t = x
		x = y
		y = t
		z = -z
	}
	// Now we are in tetrahedral group, peel off the "120-ness"
	switch (which & (4 + 8)) >> 2 {
	case 0:
	case 1:
		t = x
		x = y
		y = z
		z = t
	case 2:
		t = z
		z = y
		y = x
		x = t
	}
	// Now we are in the Klein four group, peel off the "180-ness"
	switch which & (1 + 2) {
	case 0:
	case 1:
		x = -x
		y = -y
	case 2:
		y = -y
		z = -z
	case 3:
		z = -z
		x = -x
	}
	return x, y, z
}

func applyAllSymmetries(c coord) []coord { // return the orbit
	ret := make([]coord, 0)
	for j := 0; j < 2; j++ {
		for k := 0; k < 12; k++ {
			nx, ny, nz := applySymmetry(16*j+k, c.x, c.y, c.z)
			nc := coord{nx, ny, nz}
			ret = append(ret, nc)
		}
	}
	return ret
}

func computeScanners(s1 scannerT, s2 scannerT) (map[coord]int, int) {
	res := make(map[coord]int)
	//tmp := make(map[coord]coord)

	for _, v := range s1 {
		for _, w := range s2 {
			for si, u := range applyAllSymmetries(w) {
				rc := coord{x: v.x - u.x, y: v.y - u.y, z: v.z - u.z}
				res[rc]++
				if res[rc] == 12 {
					//fmt.Println(u)
					if si > 11 {
						return res, si + 4
					} else {
						return res, si
					}
				}
			}
		}
	}
	return res, -1
}

func main() {
	f, _ := os.Open("day19_input")
	//f, _ = os.Open("day19_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([]scannerT, 0)
	// skip first line
	s.Scan()
	scanner := make([]coord, 0)
	for s.Scan() {
		t := s.Text()
		if t == "" {
			inputs = append(inputs, scanner)
			s.Scan()
			scanner = make([]coord, 0)
			continue
		}
		inps := strings.Split(s.Text(), ",")
		coords := make([]int, 0)
		for _, v := range inps {
			i, _ := strconv.ParseInt(v, 10, 0)
			coords = append(coords, int(i))
		}
		c := coord{coords[0], coords[1], coords[2]}
		scanner = append(scanner, c)
	}
	inputs = append(inputs, scanner)
	fmt.Println(len(inputs))

	// Part1
	res := 0
	todo := make([]scannerT, 0)
	todo = append(todo, inputs[0])
	//fmt.Println(inputs)
	var scanners [35]coord
	done := make(map[int]bool)
	done[0] = true
	resMap := make(map[coord]bool)
	for len(todo) > 0 {
		newTodo := make([]scannerT, 0)
		for _, v := range todo {
			for _, c := range v {
				resMap[c] = true
			}
			for ii, input := range inputs {
				if done[ii] {
					continue
				}
				r, ti := computeScanners(v, input)
				for i, v := range r {
					if v >= 12 {
						for b := 0; b < len(input); b++ {
							bc := input[b]
							bx, by, bz := applySymmetry(ti, bc.x, bc.y, bc.z)
							input[b] = coord{bx + i.x, by + i.y, bz + i.z}
							scanners[ii] = i
							resMap[input[b]] = true
						}
						newTodo = append(newTodo, input)
						done[ii] = true
					}
				}
			}
		}
		todo = newTodo
	}
	res = len(resMap)
	fmt.Println(res)

	// Part2
	res = 0
	max := float64(0)
	for i, s := range scanners {
		for _, ss := range scanners[i+1:] {
			x := s.x - ss.x
			y := s.y - ss.y
			z := s.z - ss.z
			max = math.Max(max, (math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))))
		}
	}
	res = int(max)
	fmt.Println(res)
}
