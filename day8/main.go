// Advent of Code 2021, Day 8
package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 8: Seven Segment Search
// Part 1 answer: 318
// Part 2 answer: 996280
func main() {
	fmt.Println("Advent of Code 2021, Day 8")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	// Count the number of entries on the output side for each length
	lenCount := make(map[int]int)
	for _, line := range entries {
		output := strings.Split(line, " | ")[1]
		for _, s := range strings.Split(output, " ") {
			lenCount[len(s)]++
		}
	}
	// Add up 1s (len 2), 4s (len 4), 7s (len 3), and 8s (len 7)
	return lenCount[2] + lenCount[4] + lenCount[3] + lenCount[7]
}

func part2(entries []string) int {
	var sum int
	for _, line := range entries {
		// Index is the digit, value is the string representing that digit
		var digitMap [10]string
		parts := strings.Split(line, " | ")
		signals := sortPatterns(strings.Split(parts[0], " "))
		output := sortPatterns(strings.Split(parts[1], " "))
		// Find the ones we know for sure: 1, 7, 4, 8
		for _, s := range signals {
			switch len(s) {
			case 2:
				digitMap[1] = s
			case 3:
				digitMap[7] = s
			case 4:
				digitMap[4] = s
			case 7:
				digitMap[8] = s
			}
		}
		// Now we can use deduction to find 3 and 9
		for _, s := range signals {
			// To find 3, the only len(5) that has all the segments of 7
			if len(s) == 5 && hasSegments(s, digitMap[7]) {
				digitMap[3] = s
			} else
			// To find 9, the only len(6) that has all the segments of 4
			if len(s) == 6 && hasSegments(s, digitMap[4]) {
				digitMap[9] = s
			}
		}
		// Now we can find Segment E, by diffing 9 from 8
		segE := diff(digitMap[8], digitMap[9])
		// And now that we know that, we can find the number 2 and 5
		for _, s := range signals {
			if len(s) == 5 && s != digitMap[3] {
				if !strings.ContainsRune(s, segE) {
					digitMap[5] = s
				} else {
					digitMap[2] = s
				}
			}
		}
		// Now find segment C, by diffing 9 from 5
		segC := diff(digitMap[9], digitMap[5])
		// And now we can differentiate 0 from 6
		for _, s := range signals {
			if len(s) == 6 && s != digitMap[9] {
				if !strings.ContainsRune(s, segC) {
					digitMap[6] = s
				} else {
					digitMap[0] = s
				}
			}
		}
		// Got it all
		sum += findDigit(digitMap, output[0]) * 1000
		sum += findDigit(digitMap, output[1]) * 100
		sum += findDigit(digitMap, output[2]) * 10
		sum += findDigit(digitMap, output[3])
	}
	return sum
}

// Returns true if all of the segments of s2 are contained in s1
func hasSegments(s1, s2 string) bool {
	for _, r := range s2 {
		if !strings.ContainsRune(s1, r) {
			return false
		}
	}
	return true
}

// Returns the one character that is in s1 but not s2
func diff(s1, s2 string) rune {
	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) {
			return r
		}
	}
	panic("oh noes")
}

// Returns the digit that the string represents
func findDigit(digitMap [10]string, s string) int {
	for i, digitStr := range digitMap {
		if s == sorted(digitStr) {
			return i
		}
	}
	panic("whoops")
}

// Sort each individual element in character order
func sortPatterns(input []string) []string {
	output := make([]string, 0, len(input))
	for _, s := range input {
		output = append(output, sorted(s))
	}
	return output
}

// Quick and dirty sort the characters in a string
func sorted(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
