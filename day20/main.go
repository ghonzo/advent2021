// Advent of Code 2021, Day 20
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 20: Trench Map
// Part 1 answer: 4917
// Part 2 answer: 16389
func main() {
	fmt.Println("Advent of Code 2021, Day 20")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	algorithm := []byte(entries[0])
	var grid common.Grid = common.ArraysGridFromLines(pad(entries[2:], 6))
	grid = cycle(grid, algorithm)
	grid = cycle(grid, algorithm)
	count := 0
	for p := range grid.AllPoints() {
		if grid.Get(p) == '#' {
			count++
		}
	}
	return count
}

func part2(entries []string) int {
	algorithm := []byte(entries[0])
	var grid common.Grid = common.ArraysGridFromLines(pad(entries[2:], 100))
	for i := 0; i < 50; i++ {
		grid = cycle(grid, algorithm)
	}
	count := 0
	for p := range grid.AllPoints() {
		if grid.Get(p) == '#' {
			count++
		}
	}
	return count
}

func cycle(grid common.Grid, algorithm []byte) common.Grid {
	newGrid := grid.Clone()
	for p := range grid.AllPoints() {
		val := 0
		for _, delta := range []common.Point{common.UL, common.U, common.UR, common.L, {}, common.R, common.DL, common.D, common.DR} {
			val <<= 1
			otherVal, ok := grid.CheckedGet(p.Add(delta))
			// This is what handles the infinte canvas. If it's out of bounds, then just assume the same value as the pixel.
			if !ok {
				otherVal = grid.Get(p)
			}
			if otherVal == '#' {
				val++
			}
		}
		newGrid.Set(p, algorithm[val])
	}
	return newGrid
}

// Add count '.' characters to the border of the string matrix
func pad(entries []string, count int) []string {
	padded := make([]string, len(entries)+count*2)
	blankLine := strings.Repeat(".", len(entries[0])+count*2)
	sidePad := strings.Repeat(".", count)
	for i := 0; i < count; i++ {
		padded[i] = blankLine
		padded[i+count+len(entries)] = blankLine
	}
	for i, s := range entries {
		padded[i+count] = sidePad + s + sidePad
	}
	return padded
}
