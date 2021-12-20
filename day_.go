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

func main() {
	f, _ := os.Open("day_input")
	//f, _ = os.Open("day_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([]int, 0)
	for s.Scan() {
		inps := strings.Split(s.Text(), ",")
		for _, v := range inps {
			i, _ := strconv.ParseInt(v, 10, 0)
			inputs = append(inputs, int(i))
		}
	}

	// Part1
	res := 0

	fmt.Println(res)

	// Part2
	res = 0
	fmt.Println(res)
}
