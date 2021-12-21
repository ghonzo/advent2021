// Advent of Code 2021, Day 5
package main

import (
	"fmt"
	"regexp"

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
	r := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	for _, line := range entries {
		groups := r.FindStringSubmatch(line)
		x0 := common.Atoi(groups[1])
		y0 := common.Atoi(groups[2])
		x1 := common.Atoi(groups[3])
		y1 := common.Atoi(groups[4])
		xd := common.Sgn(x1 - x0)
		yd := common.Sgn(y1 - y0)
		// Discard diagonal lines
		if xd != 0 && yd != 0 {
			continue
		}
		for x, y := x0, y0; y != y1 || x != x1; x, y = x+xd, y+yd {
			grid[common.NewPoint(x, y)]++
		}
		grid[common.NewPoint(x1, y1)]++
	}
	var sum int
	for _, v := range grid {
		if v > 1 {
			sum++
		}
	}
	return sum
}

func part2(entries []string) int {
	grid := make(map[common.Point]int)
	r := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	for _, line := range entries {
		groups := r.FindStringSubmatch(line)
		x0 := common.Atoi(groups[1])
		y0 := common.Atoi(groups[2])
		x1 := common.Atoi(groups[3])
		y1 := common.Atoi(groups[4])
		xd := common.Sgn(x1 - x0)
		yd := common.Sgn(y1 - y0)
		for x, y := x0, y0; y != y1 || x != x1; x, y = x+xd, y+yd {
			grid[common.NewPoint(x, y)]++
		}
		grid[common.NewPoint(x1, y1)]++
	}
	var sum int
	for _, v := range grid {
		if v > 1 {
			sum++
		}
	}
	return sum
}
