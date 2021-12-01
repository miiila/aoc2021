package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    f, _ := os.Open("day1_input")
    s := bufio.NewScanner(f)
    var res int = 0
    measurements := make([]int64, 0)
    for s.Scan() {
        t, _ := strconv.ParseInt(s.Text(),10, 0)
        measurements = append(measurements, t)
    }

    var prev int64 = 10000
    // Part1
    for _,v := range measurements{
        if v > prev {
            res++
        }
        prev = v
    }
    fmt.Println(res)
    // Part2
    var prevWindow int64 = 10000
    var window int64
    res = 0
    for i,v := range measurements{
        if i < len(measurements) - 2 {
            window = v + measurements[i+1] + measurements[i+2]
        }
        if window > prevWindow {
            res++
        }
        prevWindow = window
    }


    fmt.Println(res)
}
