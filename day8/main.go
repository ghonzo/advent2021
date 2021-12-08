// Advent of Code 2021, Day 8
package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 8:
// Part 1 answer:
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2021, Day 8")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	lenCount := make(map[int]int)
	for _, line := range entries {
		output := strings.Split(line, " | ")[1]
		for _, s := range strings.Split(output, " ") {
			lenCount[len(s)]++
		}
	}
	return lenCount[2] + lenCount[4] + lenCount[3] + lenCount[7]
}

func part2(entries []string) int {
	var sum int
	for _, line := range entries {
		var digitMap [10]string
		parts := strings.Split(line, " | ")
		output := strings.Split(parts[1], " ")
		digits := append(strings.Split(parts[0], " "), output...)
		// Find the ones we know for sure
		for _, s := range digits {
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
		// Now find other digits
		for _, s := range digits {
			// To find 3, the only len(5) that has all the segments of 7
			if len(s) == 5 && hasSegments(s, digitMap[7]) {
				digitMap[3] = s
				continue
			}
			// To find 9, the only len(6) that has all the segments of 4
			if len(s) == 6 && hasSegments(s, digitMap[4]) {
				digitMap[9] = s
			}
		}
		// Now we can find Segment E, by diffing 9 from 8
		segE := diff(digitMap[8], digitMap[9])
		// And now that we know that, we can find the number 2 and 5
		for _, s := range digits {
			if len(s) == 5 && sorted(s) != sorted(digitMap[3]) {
				if strings.IndexRune(s, segE) == -1 {
					digitMap[5] = s
				} else {
					digitMap[2] = s
				}
			}
		}
		// Now find segment C, by diffing 9 from 5
		segC := diff(digitMap[9], digitMap[5])
		// And now we can differentiate 0 from 6
		for _, s := range digits {
			if len(s) == 6 && sorted(s) != sorted(digitMap[9]) {
				if strings.IndexRune(s, segC) == -1 {
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

func hasSegments(s1, s2 string) bool {
	for _, r := range s2 {
		if !strings.ContainsRune(s1, r) {
			return false
		}
	}
	return true
}

func diff(s1, s2 string) rune {
	for _, r := range s1 {
		if strings.IndexRune(s2, r) == -1 {
			return r
		}
	}
	panic("oh noes")
}

func findDigit(digitMap [10]string, s string) int {
	for i, digitStr := range digitMap {
		if sorted(s) == sorted(digitStr) {
			return i
		}
	}
	panic("whoops")
}

func sorted(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
