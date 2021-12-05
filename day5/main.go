// Advent of Code 2021, Day 5
package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ghonzo/advent2021/common"
)

// Day 5: Hydrothermal Venture
// Part 1 answer: 5294
// Part 2 answer: 21698
func main() {
	fmt.Println("Advent of Code 2021, Day 5")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	grid := make(map[common.Point]int)
	r := regexp.MustCompile("(\\d+),(\\d+) -> (\\d+),(\\d+)")
	for _, line := range entries {
		groups := r.FindStringSubmatch(line)
		x0 := atoi(groups[1])
		y0 := atoi(groups[2])
		x1 := atoi(groups[3])
		y1 := atoi(groups[4])
		if x0 == x1 {
			if y0 > y1 {
				y0, y1 = y1, y0
			}
			for y := y0; y <= y1; y++ {
				p := common.NewPoint(x0, y)
				grid[p] = grid[p] + 1
			}
		} else if y0 == y1 {
			if x0 > x1 {
				x0, x1 = x1, x0
			}
			for x := x0; x <= x1; x++ {
				p := common.NewPoint(x, y0)
				grid[p] = grid[p] + 1
			}
		}
	}
	var sum int
	for _, v := range grid {
		if v > 1 {
			sum++
		}
	}
	return sum
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	}
	return 0
}

func part2(entries []string) int {
	grid := make(map[common.Point]int)
	r := regexp.MustCompile("(\\d+),(\\d+) -> (\\d+),(\\d+)")
	for _, line := range entries {
		groups := r.FindStringSubmatch(line)
		x0 := atoi(groups[1])
		y0 := atoi(groups[2])
		x1 := atoi(groups[3])
		y1 := atoi(groups[4])
		xd := sgn(x1 - x0)
		yd := sgn(y1 - y0)
		for x, y := x0, y0; y != y1 || x != x1; {
			p := common.NewPoint(x, y)
			grid[p] = grid[p] + 1
			x += xd
			y += yd
		}
		p := common.NewPoint(x1, y1)
		grid[p] = grid[p] + 1
	}
	var sum int
	for _, v := range grid {
		if v > 1 {
			sum++
		}
	}
	return sum
}
