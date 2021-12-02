package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

type position struct {
    horizontal int
    depth int
}

type instruction struct {
    direction string
    value int
}

func main() {
    f, _ := os.Open("day2_input")
    s := bufio.NewScanner(f)
    instructions := make([]instruction, 0)
    for s.Scan() {
        t := s.Text()
        vals := strings.Split(t, " ")
        i, _ := strconv.ParseInt(vals[1],10,0)
        instructions = append(instructions, instruction{vals[0], int(i)})
    }


    // Part1
    var res position
    for _, i := range instructions {
        switch i.direction {
        case "forward":
            res.horizontal += i.value
        case "down":
            res.depth += i.value
        case "up":
            res.depth -= i.value
        }
    }

    fmt.Println(res.depth*res.horizontal)

    // Part2
    type positionAim struct {
        horizontal int
        depth int
        aim int
    }

    var res2 positionAim
    for _, i := range instructions {
        switch i.direction {
        case "forward":
            res2.horizontal += i.value
            res2.depth += res2.aim * i.value
        case "down":
            res2.aim += i.value
        case "up":
            res2.aim -= i.value
        }
    }

    fmt.Println(res2.depth*res.horizontal)

}
