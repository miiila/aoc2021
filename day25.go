package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("day25_input")
	//f, _ = os.Open("day25_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([][]string, 0)
	for s.Scan() {
		row := make([]string, 0)
		for _, v := range s.Text() {
			if v == '.' {
				row = append(row, "")
			} else {
				row = append(row, string(v))
			}
		}
		inputs = append(inputs, row)
	}

	// Part1
	res := 0
	moved := true
	for moved {
		newInputs := make([][]string, 0)
		res++
		moved = false
		for i := 0; i < len(inputs); i++ {
			newInputs = append(newInputs, make([]string, len(inputs[0])))
			for j := 0; j < len(inputs[0]); j++ {
				if inputs[i][j] == ">" {
					curJ := j + 1
					if curJ == len(inputs[0]) {
						curJ = 0
					}
					if inputs[i][curJ] == "" {
						newInputs[i][curJ] = ">"
						newInputs[i][j] = ""
						moved = true
					} else {
						newInputs[i][j] = ">"
					}
				}
			}
		}
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(inputs[0]); j++ {
				if inputs[i][j] == "v" {
					curI := i + 1
					if curI == len(inputs) {
						curI = 0
					}
					if inputs[curI][j] != "v" && newInputs[curI][j] != ">" {
						newInputs[curI][j] = "v"
						newInputs[i][j] = ""
						moved = true
					} else {
						newInputs[i][j] = "v"
					}
				}
			}
		}
		inputs = newInputs
	}
	fmt.Println(res)

	// Part2
	res = 0
	fmt.Println(res)
}

func printInputs(inputs [][]string) {
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[0]); j++ {
			if inputs[i][j] == "" {
				fmt.Print(".")
			} else {
				fmt.Print(inputs[i][j])
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
