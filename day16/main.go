// Advent of Code 2021, Day 16
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 16: Packet Decoder
// Part 1 answer: 986
// Part 2 answer: 18234816469452
func main() {
	fmt.Println("Advent of Code 2021, Day 16")
	buf, _ := ioutil.ReadFile("input.txt")
	transmission := string(buf)
	fmt.Printf("Part 1: %d \n", part1(transmission))
	fmt.Printf("Part 2: %d \n", part2(transmission))
}

type packet struct {
	version int
	id      int
	// Later figure out how to do separate subtypes ... or not
	value int
	subs  []packet
}

func part1(transmission string) int {
	bin := convertToBin(transmission)
	return sumVersions(readRootPacket(bin))
}

func convertToBin(hex string) string {
	// We could use big.Int, but it doesn't handle leading zeros the way we want it
	hexMap := make(map[string]string)
	// Create the lookup table
	for i := 0; i < 16; i++ {
		hexMap[fmt.Sprintf("%X", i)] = fmt.Sprintf("%04b", i)
	}
	var sb strings.Builder
	for _, r := range hex {
		sb.WriteString(hexMap[string(r)])
	}
	return sb.String()
}

func sumVersions(p packet) int {
	sum := p.version
	for _, sp := range p.subs {
		sum += sumVersions(sp)
	}
	return sum
}

func readRootPacket(s string) packet {
	index := 0
	return readSubPacket(s, &index)
}

func readSubPacket(s string, p *int) packet {
	version := read(s, p, 3)
	id := read(s, p, 3)
	packet := packet{version: version, id: id}
	switch id {
	case 4:
		packet.value = readLiteralValue(s, p)
	default:
		packet.subs = readOperator(s, p)
	}
	return packet
}

func read(s string, p *int, n int) int {
	substr := s[*p : *p+n]
	*p += n
	i, _ := strconv.ParseInt(substr, 2, 0)
	return int(i)
}

func readLiteralValue(s string, p *int) int {
	var value int
	for {
		prefix := read(s, p, 1)
		value = value<<4 + read(s, p, 4)
		if prefix == 0 {
			return value
		}
	}
}

func readOperator(s string, p *int) []packet {
	lenType := read(s, p, 1)
	var packets []packet
	switch lenType {
	case 0:
		subpacketLen := read(s, p, 15)
		spStartIndex := *p
		for *p < spStartIndex+subpacketLen {
			packets = append(packets, readSubPacket(s, p))
		}
	case 1:
		numSubPackets := read(s, p, 11)
		for n := numSubPackets; n > 0; n-- {
			packets = append(packets, readSubPacket(s, p))
		}
	}
	return packets
}

func part2(transmission string) int {
	bin := convertToBin(transmission)
	return evaluate(readRootPacket(bin))
}

func evaluate(p packet) int {
	switch p.id {
	case 0:
		// sum
		return reduce(p.subs, 0, func(a, b int) int {
			return a + b
		})
	case 1:
		// product
		return reduce(p.subs, 1, func(a, b int) int {
			return a * b
		})
	case 2:
		// minimum
		return reduce(p.subs, math.MaxInt, func(a, b int) int {
			return common.Min(a, b)
		})
	case 3:
		// maximum
		return reduce(p.subs, 0, func(a, b int) int {
			return common.Max(a, b)
		})
	case 4:
		// literal
		return p.value
	case 5:
		// greater than
		return btoi(evaluate(p.subs[0]) > evaluate(p.subs[1]))
	case 6:
		// less than
		return btoi(evaluate(p.subs[0]) < evaluate(p.subs[1]))
	case 7:
		// equal to
		return btoi(evaluate(p.subs[0]) == evaluate(p.subs[1]))
	}
	panic("bad news man")
}

func btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func reduce(ps []packet, seed int, reduceFunc func(a, b int) int) int {
	v := seed
	for _, p := range ps {
		v = reduceFunc(v, evaluate(p))
	}
	return v
}
