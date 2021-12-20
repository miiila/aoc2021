package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"strconv"
	//"math"
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
	f, _ = os.Open("day19_input_test")
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

	// Part1
	res := 0
	//c := coord{1, 2, 3}
	//fmt.Println(inputs)
	//for si := 0; si < len(inputs)-1; si++ {
	//for sj := si + 1; sj < len(inputs); sj++ {
	//r, ti := computeScanners(inputs[si], inputs[sj])
	////fmt.Println(len(r))
	//for i, v := range r {
	//if v == 12 {
	//f := inputs[sj][0]
	//fmt.Println("Inputs[sj][0]", sj, f)
	//fmt.Println(applySymmetry(ti, f.x, f.y, f.z))
	////for b := 0; b < len(inputs[sj]); b++ {
	////bc := inputs[sj][b]
	////bx, by, bz := applySymmetry(ti, bc.x, bc.y, bc.z)
	////fmt.Println(bc, bx, by, bz)
	////inputs[sj][b] = coord{bx, by, bz}
	////}
	//fmt.Println(i, v, si, sj, ti)
	//}
	//}
	//}
	//}
	r, ti := computeScanners(inputs[0], inputs[1])
	for i, v := range r {
		if v == 12 {
			//f := inputs[sj][0]
			//fmt.Println("Inputs[sj][0]", sj, f)
			//fmt.Println(applySymmetry(ti, f.x, f.y, f.z))
			fmt.Println(i, v, ti)
		}
	}
	for b := 0; b < len(inputs[1]); b++ {
		bc := inputs[1][b]
		bx, by, bz := applySymmetry(ti, bc.x, bc.y, bc.z)
		//fmt.Println(bc, bx, by, bz)
		inputs[1][b] = coord{bx, by, bz}
	}
	r, ti = computeScanners(inputs[1], inputs[4])
	for i, v := range r {
		if v == 12 {
			//f := inputs[sj][0]
			//fmt.Println("Inputs[sj][0]", sj, f)
			//fmt.Println(applySymmetry(ti, f.x, f.y, f.z))
			fmt.Println(i, v, ti)
		}
	}

	fmt.Println(applyAllSymmetries(coord{1, 2, 3}))

	fmt.Println(res)

	// Part2
	res = 0
	fmt.Println(res)
}
