// Advent of Code 2021, Day 7
package main

import (
	"fmt"
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
	mm := new(common.MaxMin)
	crabs := make([]int, 0)
	for _, ps := range strings.Split(s, ",") {
		p := common.Atoi(ps)
		crabs = append(crabs, p)
		mm.Accept(p)
	}
	costMM := new(common.MaxMin)
	for meet := mm.Min; meet <= mm.Max; meet++ {
		var cost int
		for _, p := range crabs {
			cost += common.Abs(p - meet)
		}
		costMM.Accept(cost)
	}
	return costMM.Min
}

func part2(s string) int {
	mm := new(common.MaxMin)
	crabs := make([]int, 0)
	for _, ps := range strings.Split(s, ",") {
		p := common.Atoi(ps)
		crabs = append(crabs, p)
		mm.Accept(p)
	}
	costMM := new(common.MaxMin)
	for meet := mm.Min; meet <= mm.Max; meet++ {
		var cost int
		for _, p := range crabs {
			cost += totalFuel(common.Abs(p - meet))
		}
		costMM.Accept(cost)
	}
	return costMM.Min
}

func totalFuel(n int) int {
	// The formula for the first n postive integers is n(n+1)/2
	return n * (n + 1) / 2
}
