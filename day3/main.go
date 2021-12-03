// Advent of Code 2021, Day 3
package main

import (
	"fmt"
	"strconv"

	"github.com/ghonzo/advent2021/common"
)

// Day 3: Binary Diagnostic
// Part 1 answer: 749376
// Part 2 answer: 2372923
func main() {
	fmt.Println("Advent of Code 2021, Day 3")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries, 12))
	fmt.Printf("Part 2: %d \n", part2(entries, 12))
}

func part1(entries []string, bits int) int {
	b := make([]int, bits)
	for _, e := range entries {
		for p, ch := range e {
			if ch == '1' {
				b[p]++
			}
		}
	}
	var sum1, sum2 int
	for i := 0; i < bits; i++ {
		sum1 = sum1 << 1
		sum2 = sum2 << 1
		if b[i] > len(entries)/2 {
			sum1++
		} else {
			sum2++
		}
	}
	return sum1 * sum2
}

func part2(entries []string, bits int) int64 {
	e2 := make([]string, len(entries))
	copy(e2, entries)
	for i := 0; len(e2) > 1; i++ {
		e2 = whittle(e2, bits, i, true)
	}
	e3 := make([]string, len(entries))
	copy(e3, entries)
	for i := 0; len(e3) > 1; i++ {
		e3 = whittle(e3, bits, i, false)
	}
	ogr, _ := strconv.ParseInt(e2[0], 2, 64)
	scrubber, _ := strconv.ParseInt(e3[0], 2, 64)
	return ogr * scrubber
}

func whittle(entries []string, bits int, pos int, b bool) []string {
	originalLen := len(entries)
	sum := 0
	for _, e := range entries {
		if e[pos] == '1' {
			sum++
		}
	}
	for i := len(entries) - 1; i >= 0; i-- {
		if b {
			if (sum*2 >= originalLen && entries[i][pos] == '0') || (sum*2 < originalLen && entries[i][pos] == '1') {
				entries = append(entries[:i], entries[i+1:]...)
			}
		} else {
			if (sum*2 < originalLen && entries[i][pos] == '0') || (sum*2 >= originalLen && entries[i][pos] == '1') {
				entries = append(entries[:i], entries[i+1:]...)
			}
		}
	}
	return entries
}
