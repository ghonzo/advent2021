// Advent of Code 2021, Day 13
package main

import (
	"fmt"
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
	points, folds := parseInstructions(entries)
	return len(foldPoints(points, folds[0]))
}

func parseInstructions(entries []string) (pointset, []fold) {
	points := make(pointset)
	var i int
	var s string
	for i, s = range entries {
		if len(s) == 0 {
			break
		}
		parts := strings.Split(s, ",")
		points[common.NewPoint(common.Atoi(parts[0]), common.Atoi(parts[1]))] = true
	}
	var folds []fold
	for _, s = range entries[i+1:] {
		eqIndex := strings.IndexRune(s, '=')
		dim := s[eqIndex-1 : eqIndex]
		mag := common.Atoi(s[eqIndex+1:])
		folds = append(folds, fold{dim, mag})
	}
	return points, folds
}

type pointset map[common.Point]bool
type fold struct {
	dim string
	mag int
}

func (f fold) transform(p common.Point) common.Point {
	if f.dim == "x" {
		x := p.X()
		if x > f.mag {
			x = 2*f.mag - x
		}
		return common.NewPoint(x, p.Y())
	}
	y := p.Y()
	if y > f.mag {
		y = 2*f.mag - y
	}
	return common.NewPoint(p.X(), y)
}

func part2(entries []string) {
	points, folds := parseInstructions(entries)
	for _, f := range folds {
		points = foldPoints(points, f)
	}
	mmX := new(common.MaxMin)
	mmY := new(common.MaxMin)
	for p := range points {
		mmX.Accept(p.X())
		mmY.Accept(p.Y())
	}
	for y := 0; y <= mmY.Max; y++ {
		fmt.Println()
		for x := 0; x <= mmX.Max; x++ {
			if points[common.NewPoint(x, y)] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
	}
}

func foldPoints(points pointset, f fold) pointset {
	newPoints := make(pointset)
	for p := range points {
		newPoints[f.transform(p)] = true
	}
	return newPoints
}
