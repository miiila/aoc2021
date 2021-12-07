package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("day7_input")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([]int, 0)
	for s.Scan() {
		for _, v := range strings.Split(s.Text(), ",") {
			i, _ := strconv.ParseInt(v, 10, 0)
			inputs = append(inputs, int(i))
		}
	}

	//inputs = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	// Part1
	res := make(map[int]int)
	for _, v := range inputs {
		if res[v] > 0 {
			continue
		}
		for j := 0; j < len(inputs); j++ {
			res[v] = res[v] + int(math.Abs(float64(v-inputs[j])))
		}
	}

	minF := math.MaxInt
	for _, v := range res {
		if v < minF {
			minF = v
		}
	}
	fmt.Println(minF)

	// Part2
	res = make(map[int]int)
	for _, v := range inputs {
		if res[v] > 0 {
			continue
		}
		for j := 0; j < len(inputs); j++ {
			d := int(math.Abs(float64(v - inputs[j])))
			res[v] = res[v] + (d * (d + 1) / 2)
		}
	}

	minF = math.MaxInt
	for _, v := range res {
		if v < minF {
			minF = v
		}
	}
	fmt.Println(minF)
}
