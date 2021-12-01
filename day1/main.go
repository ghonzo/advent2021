// Advent of Code 2021, Day 1
package main

import (
	"fmt"

	"github.com/ghonzo/advent2021/common"
)

// Day 1: Sonar Sweep
// Part 1 answer: 1527
// Part 2 answer: 1575
func main() {
	fmt.Println("Advent of Code 2021, Day 1")
	entries := common.ReadIntsFromFile("input.txt")
	fmt.Printf("Part 1: %d increases\n", part1(entries))
	fmt.Printf("Part 2: %d increases\n", part2(entries))
}

func part1(entries []int) int {
	// Find the number of increases
	increases := 0
	last := 0
	for _, e := range entries {
		if last != 0 && e > last {
			increases++
		}
		last = e
	}
	return increases
}

func part2(entries []int) int {
	// Find the number of increases
	increases := 0
	last := 0
	for i, e := range entries[2:] {
		sum := entries[i] + entries[i+1] + e
		if last != 0 && sum > last {
			increases++
		}
		last = sum
	}
	return increases
}
