package main

import (
	"fmt"
)

type dice struct {
	state int
	rolls int
}

func (d *dice) roll() int {
	d.state++
	d.rolls++
	if d.state > 100 {
		d.state = 1
	}

	return d.state
}

func (d *dice) roll3() int {
	k := 0
	for i := 0; i < 3; i++ {
		n := d.roll()
		k = k + n
	}

	return k
}

type finalres struct {
	p1 int
	p2 int
}

func main() {
	// Part1
	res := 0
	d := dice{0, 0}
	p1s, p2s, p1p, p2p, i := 0, 0, 8, 6, 0
	for p1s < 1000 && p2s < 1000 {
		i++
		i = i % 2
		if i == 1 {
			p1p = (p1p + d.roll3()) % 10
			if p1p == 0 {
				p1p = 10
			}
			p1s = p1s + p1p
		} else {
			p2p = (p2p + d.roll3()) % 10
			if p2p == 0 {
				p2p = 10
			}
			p2s = p2s + p2p
		}
	}

	res = p2s * d.rolls
	fmt.Println(res)

	// Part2
	res = 0
  fres := finalres{p1: 0, p2: 0}
	st := getDiceStates()
	diracDice(8, 0, 6, 0, 0, &st, &fres, 1)
	fmt.Println(fres)
}

func diracDice(p1p int, p1s int, p2p int, p2s int, turn int, diceStates *map[int]int, res *finalres, mult int) {
	if p1s >= 21 {
		res.p1 = res.p1 + mult
    return
	}
  if p2s >= 21 {
		res.p2 = res.p2 + mult
    return
	}
		for k, v := range *diceStates {
			if turn == 0 {
				np1p := (p1p + k) % 10
				if np1p == 0 {
					np1p = 10
				}
				np1s := p1s + np1p
				diracDice(np1p, np1s, p2p, p2s, 1, diceStates, res, v*mult)
			} else {
        np2p := (p2p + k) % 10
				if np2p == 0 {
					np2p = 10
				}
        np2s := p2s + np2p
				diracDice(p1p, p1s, np2p, np2s, 0, diceStates, res, v*mult)
			}
		}
  return
}

func getDiceStates() map[int]int {
	res := make(map[int]int, 0)
	for i := 3; i > 0; i-- {
		for j := 3; j > 0; j-- {
			for k := 3; k > 0; k-- {
				res[i+j+k]++
			}
		}
	}

	return res
}
