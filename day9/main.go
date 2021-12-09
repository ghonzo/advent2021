// Advent of Code 2021, Day 9
package main

import (
	"fmt"
	"sort"

	"github.com/ghonzo/advent2021/common"
)

// Day 9: Smoke Basin
// Part 1 answer: 417
// Part 2 answer: 1148965
func main() {
	fmt.Println("Advent of Code 2021, Day 9")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	var sum int
	grid := common.ArraysGridFromLines(entries)
	for pt := range grid.AllPoints() {
		val := grid.Get(pt)
		if v, ok := grid.CheckedGet(pt.Add(common.N)); ok && v <= val {
			continue
		}
		if v, ok := grid.CheckedGet(pt.Add(common.E)); ok && v <= val {
			continue
		}
		if v, ok := grid.CheckedGet(pt.Add(common.S)); ok && v <= val {
			continue
		}
		if v, ok := grid.CheckedGet(pt.Add(common.W)); ok && v <= val {
			continue
		}
		// Must be the lowest
		sum += int(val - '0' + 1)
	}
	return sum
}

func part2(entries []string) int {
	grid := common.ArraysGridFromLines(entries)
	basins := make([]map[common.Point]bool, 0)
	foundPoints := make(map[common.Point]bool)
	for pt := range grid.AllPoints() {
		if foundPoints[pt] {
			continue
		}
		foundPoints[pt] = true
		if grid.Get(pt) == '9' {
			continue
		}
		basin := make(map[common.Point]bool)
		basin[pt] = true
		expandBasin(basin, grid)
		basins = append(basins, basin)
		for k, _ := range basin {
			foundPoints[k] = true
		}
	}
	sizes := make([]int, len(basins))
	for _, b := range basins {
		sizes = append(sizes, len(b))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func expandBasin(b map[common.Point]bool, g common.Grid) {
	size0 := len(b)
	pointsToAdd := make([]common.Point, 0)
	for pt, _ := range b {
		if v, ok := g.CheckedGet(pt.Add(common.N)); ok && v < '9' {
			pointsToAdd = append(pointsToAdd, pt.Add(common.N))
		}
		if v, ok := g.CheckedGet(pt.Add(common.E)); ok && v < '9' {
			pointsToAdd = append(pointsToAdd, pt.Add(common.E))
		}
		if v, ok := g.CheckedGet(pt.Add(common.S)); ok && v < '9' {
			pointsToAdd = append(pointsToAdd, pt.Add(common.S))
		}
		if v, ok := g.CheckedGet(pt.Add(common.W)); ok && v < '9' {
			pointsToAdd = append(pointsToAdd, pt.Add(common.W))
		}
	}
	for _, pt := range pointsToAdd {
		b[pt] = true
	}
	if len(b) > size0 {
		expandBasin(b, g)
	}
}
