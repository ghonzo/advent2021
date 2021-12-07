// Advent of Code 2021, Day 7
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 7: The Treachery of Whales
// Part 1 answer: 343468
// Part 2 answer: 96086265
func main() {
	fmt.Println("Advent of Code 2021, Day 7")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries[0]))
	fmt.Printf("Part 2: %d \n", part2(entries[0]))
}

func part1(s string) int {
	var min, max int
	crabs := make([]int, 0)
	for _, ps := range strings.Split(s, ",") {
		p := atoi(ps)
		crabs = append(crabs, p)
		if p > max {
			max = p
		}
		if p < min || min == 0 {
			min = p
		}
	}
	var mincost int
	for meet := min; meet <= max; meet++ {
		var cost int
		for _, p := range crabs {
			cost += common.Abs(p - meet)
		}
		if cost < mincost || mincost == 0 {
			mincost = cost
		}
	}
	return mincost
}

func part2(s string) int {
	var min, max int
	crabs := make([]int, 0)
	for _, ps := range strings.Split(s, ",") {
		p := atoi(ps)
		crabs = append(crabs, p)
		if p > max {
			max = p
		}
		if p < min || min == 0 {
			min = p
		}
	}
	var mincost int
	for meet := min; meet <= max; meet++ {
		var cost int
		for _, p := range crabs {
			cost += totalFuel(common.Abs(p - meet))
		}
		if cost < mincost || mincost == 0 {
			mincost = cost
		}
	}
	return mincost
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func totalFuel(n int) int {
	// The formula for the first n postive integers is n(n+1)/2
	return n * (n + 1) / 2
}
