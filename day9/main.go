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

type basin map[common.Point]bool

func part2(entries []string) int {
	grid := common.ArraysGridFromLines(entries)
	var basins []basin
	// Even though it's the same type as basin, we think of it differently
	foundPoints := make(map[common.Point]bool)
	for pt := range grid.AllPoints() {
		if foundPoints[pt] {
			continue
		}
		foundPoints[pt] = true
		if grid.Get(pt) == '9' {
			continue
		}
		b := make(basin)
		expandBasin(b, pt, grid)
		basins = append(basins, b)
		// Add all the points in the basin to "foundPoints"
		for k := range b {
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

func expandBasin(b basin, pt common.Point, grid common.Grid) {
	if v, ok := grid.CheckedGet(pt); ok && v < '9' && !b[pt] {
		b[pt] = true
		expandBasin(b, pt.Add(common.N), grid)
		expandBasin(b, pt.Add(common.E), grid)
		expandBasin(b, pt.Add(common.S), grid)
		expandBasin(b, pt.Add(common.W), grid)
	}
}
