package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	//"sort"
)

type aluinstruction struct {
	inst string
	a    string
	breg string
	bint int
}

func main() {
	f, _ := os.Open("day24_input")
	//f, _ = os.Open("day24_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	inputs := make([]aluinstruction, 0)
	for s.Scan() {
		inps := strings.Split(s.Text(), " ")
		ins := aluinstruction{inst: inps[0], a: inps[1]}
		if len(inps) > 2 {
			if inps[2] == "w" || inps[2] == "x" || inps[2] == "y" || inps[2] == "z" {
				ins.breg = inps[2]
			} else {
				val, _ := strconv.ParseInt(inps[2], 10, 0)
				ins.bint = int(val)
			}
		}

		inputs = append(inputs, ins)
	}

	// Part1
	res := 0

	inputNum := "33178782929739"
	inputNum = "94399898949959"
	inputNum = "21176121611511"

	//inpNums := 0
	//for true {
	regs := make(map[string]int)
	for _, ins := range inputs {
		fmt.Println(ins, regs)
		switch ins.inst {
		case "inp":
			fmt.Println(" ")
			digit, _ := strconv.ParseInt(string(inputNum[0]), 10, 0)
			inputNum = inputNum[1:]
			regs[ins.a] = int(digit)
		case "add":
			if ins.breg == "" {
				regs[ins.a] = regs[ins.a] + ins.bint
			} else {
				regs[ins.a] = regs[ins.a] + regs[ins.breg]
			}
		case "mul":
			if ins.breg == "" {
				regs[ins.a] = regs[ins.a] * ins.bint
			} else {
				regs[ins.a] = regs[ins.a] * regs[ins.breg]
			}
		case "div":
			if ins.breg == "" {
				regs[ins.a] = regs[ins.a] / ins.bint
			} else {
				regs[ins.a] = regs[ins.a] / regs[ins.breg]
			}
		case "mod":
			if ins.breg == "" {
				regs[ins.a] = regs[ins.a] % ins.bint
			} else {
				regs[ins.a] = regs[ins.a] % regs[ins.breg]
			}
		case "eql":
			if ins.breg == "" {
				if regs[ins.a] == ins.bint {
					regs[ins.a] = 1
				} else {
					regs[ins.a] = 0
				}
			} else {
				if regs[ins.a] == regs[ins.breg] {
					regs[ins.a] = 1
				} else {
					regs[ins.a] = 0
				}
			}
		}
	}
	//if regs["z"] < 200 {
	//fmt.Println(lastInput, regs)
	//}
	if regs["z"] == 0 {
		fmt.Println("RES FOUND")
	}
	//}

	fmt.Println(res)

	// Part2
	res = 0
	fmt.Println(res)
}

func genInput() string {
	output := ""
	nums := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for len(output) < 14 {
		output = output + nums[rand.Int()%9]
	}

	return output

}
