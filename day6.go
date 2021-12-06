package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  f, _ := os.Open("day6_input")
	//f, _ := os.Open("day6_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([]int, 0)
	for s.Scan() {
		for _, v := range strings.Split(s.Text(), ",") {
			i, _ := strconv.ParseInt(v, 10, 0)
			inputs = append(inputs, int(i))
		}
	}

	// Part1
	currentGen := inputs
	for d := 0; d < 80; d++ {
		news := 0
		nextGen := make([]int, len(currentGen))
		for i, v := range currentGen {
			next := v - 1
			if next == -1 {
				next = 6
				news++
			}
			nextGen[i] = next
		}
		for n := 0; n < news; n++ {
			nextGen = append(nextGen, 8)
		}
		currentGen = nextGen
	}
	res := len(currentGen)
	fmt.Println(res)

	// Part2
	currentGen = inputs
	var days [257]int
	days[0] = len(currentGen)
	for _, v := range currentGen {
		for i := 1; i < 257-v; {
			days[v+i]++
			i = i + v
			v = 7
		}
	}
  for i:=1; i < 257;i++ {
    p := days[i]
    if p > 0 {
      for j:=i+9; j < 257; {
        days[j] = days[j]+p
        j = j+7
      }
    }
  }
  for i := 0; i <256; i++ {
    days[i+1] = days[i+1] + days[i]
  }
  if days[80] != res {
    panic("Wrong")
  }
  println(days[256])
}
