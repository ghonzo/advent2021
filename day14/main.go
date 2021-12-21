// Advent of Code 2021, Day 14
package main

import (
	"fmt"

	"github.com/ghonzo/advent2021/common"
)

// Day 14: Extended Polymerization
// Part 1 answer: 2321
// Part 2 answer: 2399822193707
func main() {
	fmt.Println("Advent of Code 2021, Day 14")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

type Pair [2]byte

type Rule struct {
	left      Pair
	generates [2]Pair
}

func part1(entries []string) int {
	return doCycles(entries, 10)
}

func part2(entries []string) int {
	return doCycles(entries, 40)
}

func doCycles(entries []string, cycles int) int {
	rules := make(map[Pair]Rule)
	for _, s := range entries[2:] {
		rule := createRule(s[0:2], s[6])
		rules[rule.left] = rule
	}
	initialTemplate := entries[0]
	// Keep track of how many Pairs
	pairCount := make(map[Pair]int)
	// And the count of each character
	charCount := make(map[byte]int)
	// Seed pairCount with the initial template
	for i := 0; i < len(initialTemplate)-1; i++ {
		charCount[initialTemplate[i]]++
		pairCount[Pair{initialTemplate[i], initialTemplate[i+1]}]++
	}
	// Don't forget the last character
	charCount[initialTemplate[len(initialTemplate)-1]]++
	// Now iterate
	for cycle := 0; cycle < cycles; cycle++ {
		nextPairCount := make(map[Pair]int)
		for p, count := range pairCount {
			// Each pair generates two more pairs with the same count
			rg := rules[p].generates
			nextPairCount[rg[0]] += count
			nextPairCount[rg[1]] += count
			// We add one character
			charCount[rg[0][1]] += count
		}
		pairCount = nextPairCount
	}
	return calculateRange(charCount)
}

func createRule(left string, right byte) Rule {
	var r Rule
	r.left = Pair{left[0], left[1]}
	// Generates two pairs
	r.generates = [2]Pair{{left[0], right}, {right, left[1]}}
	return r
}

// Range is the difference between the maximum and the minimum of the values
func calculateRange(m map[byte]int) int {
	mm := new(common.MaxMin)
	for _, v := range m {
		mm.Accept(v)
	}
	return mm.Max - mm.Min
}
