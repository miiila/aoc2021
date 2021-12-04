package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("day4_input")
	s := bufio.NewScanner(f)
	var numbers []int
	boards := make([][]int, 0)
	var board []int
	board = make([]int, 0)
	s.Scan()
	for _, v := range strings.Split(s.Text(), ",") {
		i, _ := strconv.ParseInt(v, 10, 0)
		numbers = append(numbers, int(i))
	}
	s.Scan()
	for s.Scan() {
		t := s.Text()
		if t == "" {
			boards = append(boards, board)
			board = make([]int, 0)
			continue
		}
		for _, v := range strings.Split(t, " ") {
			i, _ := strconv.ParseInt(v, 10, 0)
			if v != "" {
				board = append(board, int(i))
			}
		}
	}
	boards = append(boards, board)

	// Part1
	boards1 := make([][]int, len(boards))
	for i := range boards {
		boards1[i] = make([]int, len(boards[i]))
		copy(boards1[i], boards[i])
	}

part1:
	for _, v := range numbers {
		for _, b := range boards1 {
			for i, n := range b {
				if n == v {
					b[i] = -1
					if checkWin(b, i) {
						res := 0
						for _, k := range b {
							if k != -1 {
								res += k
							}
						}
						fmt.Println(res, n, res*n)
						break part1
					}
				}
			}
		}
	}

	// Part2
part2:
	for true {
		finalres := 0
		for _, v := range numbers {
			toRemove := make(map[int]bool, 0)
			for ib, b := range boards {
				for i, n := range b {
					if n == v {
						b[i] = -1
						if checkWin(b, i) {
							res := 0
							for _, k := range b {
								if k != -1 {
									res += k
								}
							}
							finalres = res * n
							toRemove[ib] = true
						}
					}

				}
			}
			newBoards := make([][]int, 0)
			for ib, b := range boards {
				_, in := toRemove[ib]
				if !in {
					newBoards = append(newBoards, b)
				}
			}
			boards = newBoards
			if len(boards) == 0 {
				println(finalres)
				break part2
			}
		}
	}

}
func checkWin(board []int, p int) bool {
	row := 5 * (p / 5)
	col := (p % 5)
	rowWin := true
	colWin := true
	for i := 0; i < 5; i++ {
		if board[row+i] != -1 {
			rowWin = false
		}
		if board[col+i*5] != -1 {
			colWin = false
		}
	}

	return rowWin || colWin
}
