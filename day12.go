package main

import (
	"bufio"
	"fmt"
	"os"

	//"strconv"
	//"math"
	//"sort"
	"strings"
)

func main() {
	f, _ := os.Open("day12_input")
	//f, _ = os.Open("day12_input_test")
	defer f.Close()
	s := bufio.NewScanner(f)
	caves := make(map[string][]string)
	for s.Scan() {
		spl := strings.Split(s.Text(), "-")
		caves[spl[0]] = append(caves[spl[0]], spl[1])
		caves[spl[1]] = append(caves[spl[1]], spl[0])
	}

	// Part1
	res := 0

	visited := make(map[string]bool)
	visited["start"] = true
	res = navigate("start", caves, visited)
	fmt.Println(res)

	// Part2
	res = 0
	visited = make(map[string]bool)
	visited["start"] = true
	res = navigate2("start", caves, visited, false)
	fmt.Println(res)
}

func navigate(curCave string, caves map[string][]string, visited map[string]bool) int {
	if curCave == "end" {
		return 1
	}
	toGo, ok := caves[curCave]
	if !ok {
		return 0
	}

	p := 0
	for _, c := range toGo {
		if strings.ToUpper(c) != c && visited[c] {
			continue
		}
		newVisited := make(map[string]bool)
		for k, v := range visited {
			newVisited[k] = v
		}
		if strings.ToUpper(c) != c {
			newVisited[c] = true
		}
		p = p + navigate(c, caves, newVisited)
	}

	return p

}

func navigate2(curCave string, caves map[string][]string, visited map[string]bool, usedTwice bool) int {
	if curCave == "end" {
		return 1
	}
	toGo, ok := caves[curCave]
	if !ok {
		return 0
	}

	p := 0
	for _, c := range toGo {
		if strings.ToUpper(c) != c && visited[c] && usedTwice || c == "start" {
			continue
		}
		newVisited := make(map[string]bool)
		for k, v := range visited {
			newVisited[k] = v
		}
		if strings.ToUpper(c) != c {
			newVisited[c] = true
		}
		p = p + navigate2(c, caves, newVisited, visited[c] || usedTwice)
	}

	return p

}
