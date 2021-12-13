// Advent of Code 2021, Day 13
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 13: Transparent Origami
// Part 1 answer: 710
// Part 2 answer: EPLGRULR
func main() {
	fmt.Println("Advent of Code 2021, Day 13")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	part2(entries)
	//fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	var points []common.Point
	for _, s := range entries {
		if len(s) == 0 {
			break
		}
		parts := strings.Split(s, ",")
		points = append(points, common.NewPoint(atoi(parts[0]), atoi(parts[1])))
	}
	grid := make(map[common.Point]bool)
	const xfold = 655
	for _, p := range points {
		x := p.X()
		if x > xfold {
			x = 2*xfold - x
		}
		grid[common.NewPoint(x, p.Y())] = true
	}
	return len(grid)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func part2(entries []string) {
	points := make(map[common.Point]bool)
	var i int
	var s string
	for i, s = range entries {
		if len(s) == 0 {
			break
		}
		parts := strings.Split(s, ",")
		points[common.NewPoint(atoi(parts[0]), atoi(parts[1]))] = true
	}
	for _, s = range entries[i+1:] {
		eqIndex := strings.IndexRune(s, '=')
		dim := s[eqIndex-1 : eqIndex]
		mag := atoi(s[eqIndex+1:])
		points = fold(points, dim, mag)
	}
	var maxX, maxY int
	for p := range points {
		if p.X() > maxX {
			maxX = p.X()
		}
		if p.Y() > maxY {
			maxY = p.Y()
		}
	}
	for y := 0; y <= maxY; y++ {
		fmt.Println()
		for x := 0; x <= maxX; x++ {
			if points[common.NewPoint(x, y)] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
	}
}

func fold(points map[common.Point]bool, dim string, mag int) map[common.Point]bool {
	newPoints := make(map[common.Point]bool)
	if dim == "x" {
		for p := range points {
			x := p.X()
			if x > mag {
				x = 2*mag - x
			}
			newPoints[common.NewPoint(x, p.Y())] = true
		}
	} else {
		for p := range points {
			y := p.Y()
			if y > mag {
				y = 2*mag - y
			}
			newPoints[common.NewPoint(p.X(), y)] = true
		}
	}
	return newPoints
}
