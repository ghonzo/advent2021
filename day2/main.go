// Advent of Code 2021, Day 2
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 2: Dive!
// Part 1 answer: 1840243
// Part 2 answer: 1727785422
func main() {
	fmt.Println("Advent of Code 2021, Day 2")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	// Find the final position, return x*y
	var x, y int
	for _, e := range entries {
		items := strings.Split(e, " ")
		v := common.Atoi(items[1])
		switch items[0] {
		case "forward":
			x += v
		case "down":
			y += v
		case "up":
			y -= v
		}
	}
	return x * y
}

func part2(entries []string) int {
	// Find the final position, return x*y
	var x, y, aim int
	for _, e := range entries {
		items := strings.Split(e, " ")
		v := common.Atoi(items[1])
		switch items[0] {
		case "forward":
			x += v
			y += aim * v
		case "down":
			aim += v
		case "up":
			aim -= v
		}
	}
	return x * y
}
