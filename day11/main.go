// Advent of Code 2021, Day 11
package main

import (
	"fmt"

	"github.com/ghonzo/advent2021/common"
)

// Day 11: Dumbo Octopus
// Part 1 answer: 1652
// Part 2 answer: 220
func main() {
	fmt.Println("Advent of Code 2021, Day 11")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	grid := readGridAsNumbers(entries)
	var sum int
	for step := 1; step <= 100; step++ {
		sum += cycle(grid)
	}
	return sum
}

func readGridAsNumbers(entries []string) common.Grid {
	grid := common.ArraysGridFromLines(entries)
	for pt := range grid.AllPoints() {
		grid.Set(pt, grid.Get(pt)-'0')
	}
	return grid
}

// Mutates the entries of the grid and returns the number of flashes in this cycle
func cycle(grid common.Grid) int {
	for pt := range grid.AllPoints() {
		grid.Set(pt, grid.Get(pt)+1)
	}
	flashPoints := make(map[common.Point]bool)
restart:
	for {
		for pt := range grid.AllPoints() {
			if grid.Get(pt) > 9 && !flashPoints[pt] {
				flashPoints[pt] = true
				flash(grid, pt)
				continue restart
			}
		}
		for pt := range flashPoints {
			grid.Set(pt, 0)
		}
		return len(flashPoints)
	}
}

func flash(grid common.Grid, pt common.Point) {
	for p := range pt.SurroundingPoints() {
		if v, ok := grid.CheckedGet(p); ok {
			grid.Set(p, v+1)
		}
	}
}

func part2(entries []string) int {
	grid := readGridAsNumbers(entries)
	// When this many flash in one cycle, we're done
	all := grid.Size().X() * grid.Size().Y()
	for step := 1; ; step++ {
		if cycle(grid) == all {
			return step
		}
	}
}
