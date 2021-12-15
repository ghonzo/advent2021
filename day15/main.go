// Advent of Code 2021, Day 15
package main

import (
	"container/heap"
	"fmt"

	"github.com/ghonzo/advent2021/common"
)

// Day 15: Chiton
// Part 1 answer: 702
// Part 2 answer: 2955
func main() {
	fmt.Println("Advent of Code 2021, Day 15")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	grid := readDigitGrid(entries)
	return findLowestCost(grid)
}

func part2(entries []string) int {
	grid := readDigitGrid(entries)
	return findLowestCost(makeMegaGrid(grid))
}

func findLowestCost(grid common.Grid) int {
	// mincost stores the minimum cost of all visited points
	minCost := make(map[common.Point]int)
	destination := grid.Size().Add(common.UL)
	pq := make(PriorityQueue, 0)
	pq.Push(&pointState{pt: common.NewPoint(0, 0), cost: 0})
	// heap package is what actually maintains the priority queue and what we use to push and pop
	heap.Init(&pq)
	// Use Dijkstra's algorithm (https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)
	for pq.Len() > 0 {
		ps := heap.Pop(&pq).(*pointState)
		// Did we make it?
		if ps.pt == destination {
			// Made it!
			return ps.cost
		}
		// Test each direction
		for surroundingPt := range ps.pt.SurroundingCardinals() {
			if v, ok := grid.CheckedGet(surroundingPt); ok {
				nextCost := ps.cost + int(v)
				// Is that lower than the current minCost for that point?
				if currentMinCost, ok := minCost[surroundingPt]; !ok || nextCost < currentMinCost {
					// Let's explore it
					minCost[surroundingPt] = nextCost
					pq.Push(&pointState{pt: surroundingPt, cost: nextCost})
				}
			}
		}
	}
	panic("whoops")
}

func readDigitGrid(entries []string) common.Grid {
	grid := common.ArraysGridFromLines(entries)
	common.MapGridValues(grid, func(v byte) byte {
		return v - '0'
	})
	return grid
}

type pointState struct {
	pt   common.Point
	cost int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func makeMegaGrid(grid common.Grid) common.Grid {
	size := grid.Size()
	megaGrid := common.NewArraysGrid(size.X()*5, size.Y()*5)
	for pt := range megaGrid.AllPoints() {
		divX := pt.X() / size.X()
		modX := pt.X() % size.X()
		divY := pt.Y() / size.Y()
		modY := pt.Y() % size.Y()
		baseValue := grid.Get(common.NewPoint(modX, modY))
		megaGrid.Set(pt, ((baseValue+byte(divX+divY)-1)%9)+1)
	}
	return megaGrid
}

// This whole PriorityQueue implementation was taken directly from
// https://pkg.go.dev/container/heap#example__priorityQueue

// A PriorityQueue implements heap.Interface and holds pointStates.
type PriorityQueue []*pointState

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want to order first by min cost, and then by (reverse) distance from the origin
	if pq[i].cost == pq[j].cost {
		return pq[i].pt.ManhattanDistance() > pq[j].pt.ManhattanDistance()
	}
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pointState)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
