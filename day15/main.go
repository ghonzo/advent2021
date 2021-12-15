// Advent of Code 2021, Day 15
package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/ghonzo/advent2021/common"
)

// Day 15: Chiton
// Part 1 answer: 702
// Part 2 answer: 2955
func main() {
	fmt.Println("Advent of Code 2021, Day 15")
	entries := common.ReadStringsFromFile("input.txt")
	//fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	globalMinCost = math.MaxInt
	grid := common.ArraysGridFromLines(entries)
	adjustValues(grid)
	minCost := make(map[common.Point]int)
	cost := 0
	pt := common.NewPoint(0, 0)
	// Just a signal value
	minCost[pt] = 1
	mc, ok := findMinCost(grid, minCost, cost, pt)
	if !ok {
		panic("whoops")
	}
	return mc
}

func adjustValues(grid common.Grid) {
	for pt := range grid.AllPoints() {
		grid.Set(pt, grid.Get(pt)-'0')
	}
}

var globalMinCost int

// First return value is cost, next is true if it is a valid path
func findMinCost(grid common.Grid, minCost map[common.Point]int, cost int, pt common.Point) (int, bool) {
	if cost >= globalMinCost {
		return cost, false
	}
	if pt == grid.Size().Add(common.NW) {
		if cost < globalMinCost {
			fmt.Println(cost)
			globalMinCost = cost
		}
		return cost, true
	}
	currentMin := math.MaxInt
	type pointAndCost struct {
		point common.Point
		cost  int
		score int
	}
	var nextPoints []pointAndCost
	for surroundingPt := range pt.SurroundingCardinals() {
		if v, ok := grid.CheckedGet(surroundingPt); ok {
			nextCost := cost + int(v)
			// Is that lower than the current minCost for that point?
			if currentMinCost, ok := minCost[surroundingPt]; !ok || nextCost < currentMinCost {
				minCost[surroundingPt] = nextCost
				nextPoints = append(nextPoints, pointAndCost{surroundingPt, nextCost, -surroundingPt.ManhattanDistance()*10 + nextCost})
			}
		}
	}
	// Now sort with lowest cost first
	sort.Slice(nextPoints, func(i, j int) bool {
		return nextPoints[i].score < nextPoints[j].score
	})
	// Let's see which paths we need to take
	for _, next := range nextPoints {
		if nextStepMinCost, ok := findMinCost(grid, minCost, next.cost, next.point); ok && nextStepMinCost < currentMin {
			currentMin = nextStepMinCost
		}
	}
	return currentMin, currentMin < math.MaxInt
}

func part2(entries []string) int {
	globalMinCost = math.MaxInt
	grid := common.ArraysGridFromLines(entries)
	adjustValues(grid)
	megaGrid := makeMegaGrid(grid)
	minCost := make(map[common.Point]int)
	cost := 0
	pt := common.NewPoint(0, 0)
	// Just a signal value
	minCost[pt] = 1
	mc, ok := findMinCost(megaGrid, minCost, cost, pt)
	if !ok {
		panic("whoops")
	}
	return mc
}

func makeMegaGrid(grid common.Grid) common.Grid {
	size := grid.Size()
	megaGrid := make(common.ArraysGrid, size.Y()*5)
	for row := range megaGrid {
		megaGrid[row] = make([]byte, size.X()*5)
	}
	for pt := range megaGrid.AllPoints() {
		divX := pt.X() / size.X()
		modX := pt.X() % size.X()
		divY := pt.Y() / size.Y()
		modY := pt.Y() % size.Y()
		baseValue := grid.Get(common.NewPoint(modX, modY))
		megaGrid.Set(pt, ((baseValue+byte(divX+divY)-1)%9)+1)
	}
	return &megaGrid
}
