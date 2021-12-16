package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
	//"strconv"
	//"math"
	//"sort"
)

type packet struct {
	version    int
	typeId     int
	value      int
	subPackets []packet
}

func main() {
	f, _ := os.Open("day16_input")
	defer f.Close()
	s := bufio.NewScanner(f)
	bin := ""
	for s.Scan() {
		bin = hexToBin(s.Text())
	}

	if hexToBin("38006F45291200") != "00111000000000000110111101000101001010010001001000000000" {
		panic("Parsing failed")
	}

	// Part1
	res := 0

	p := bin
	packets, _ := parseNextPacket(p)

	res = countVersions(packets)
	fmt.Println(res)

	// Part2
	res = 0

	res = solvePacket(packets)
	fmt.Println(res)
}

func hexToBin(hex string) string {
	bin := ""
	for _, v := range hex {
		i, _ := strconv.ParseInt(string(v), 16, 0)
		bin = bin + fmt.Sprintf("%04b", int(i))
	}

	return bin
}

func parseNextPacket(p string) (packet, string) {
	pk := packet{subPackets: make([]packet, 0)}
	v, _ := strconv.ParseInt(p[0:3], 2, 0)
	t, _ := strconv.ParseInt(p[3:6], 2, 0)
	p = p[6:]

	if t == 4 {
		val := ""
		i := 0
		for true {
			i++
			ns := p[0:5]
			val = val + ns[1:]
			p = p[5:]
			if ns[0] == '0' {
				break
			}
		}
		s, _ := strconv.ParseInt(val, 2, 0)

		pk.typeId = int(t)
		pk.value = int(s)
		pk.version = int(v)
		return pk, p
	}

	l := p[0]

	if l == '0' {
		length, _ := strconv.ParseInt(p[1:16], 2, 0)
		pk.typeId = int(t)
		pk.version = int(v)
		s := p[16 : 16+length]
		for len(s) > 0 {
			np, ns := parseNextPacket(s)
			pk.subPackets = append(pk.subPackets, np)
			s = ns
		}
		return pk, p[16+length:]
	} else {
		length, _ := strconv.ParseInt(p[1:12], 2, 0)
		pk.typeId = int(t)
		pk.version = int(v)
		s := p[12:]
		for i := 0; i < int(length); i++ {
			np, ns := parseNextPacket(s)
			pk.subPackets = append(pk.subPackets, np)
			s = ns
		}
		return pk, s
	}

}

func countVersions(p packet) int {
	res := p.version
	if p.typeId == 4 {
		return res
	}
	for _, v := range p.subPackets {
		res = res + countVersions(v)
	}

	return res
}

func solvePacket(p packet) int {
	if p.typeId == 4 {
		return p.value
	}
	res := 0
	if p.typeId == 0 {
		for _, v := range p.subPackets {
			res = res + solvePacket(v)
		}
		return res
	}
	if p.typeId == 1 {
		res = 1
		for _, v := range p.subPackets {
			res = res * solvePacket(v)
		}
		return res
	}
	if p.typeId == 2 {
		res = 1e16
		for _, v := range p.subPackets {
			r := solvePacket(v)
			if r < res {
				res = r
			}
		}
		return res
	}
	if p.typeId == 3 {
		res = 0
		for _, v := range p.subPackets {
			r := solvePacket(v)
			if r > res {
				res = r
			}
		}
		return res
	}
	if p.typeId == 5 {
		res = 0
		if solvePacket(p.subPackets[0]) > solvePacket(p.subPackets[1]) {
			res = 1
		}
		return res
	}
	if p.typeId == 6 {
		res = 0
		if solvePacket(p.subPackets[0]) < solvePacket(p.subPackets[1]) {
			res = 1
		}
		return res
	}
	if p.typeId == 7 {
		res = 0
		if solvePacket(p.subPackets[0]) == solvePacket(p.subPackets[1]) {
			res = 1
		}
		return res
	}

	return res
}
