// Advent of Code 2021, Day 6
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 6: Lanternfish
// Part 1 answer: 362666
// Part 2 answer: 1640526601595
func main() {
	fmt.Println("Advent of Code 2021, Day 6")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries[0]))
	fmt.Printf("Part 2: %d \n", part2(entries[0]))
}

func part1(s string) int {
	fish := make([]int, 0)
	for _, a := range strings.Split(s, ",") {
		fish = append(fish, atoi(a))
	}
	for i := 0; i < 80; i++ {
		fish = cycle(fish)
	}
	return len(fish)
}

func cycle(fish []int) []int {
	next := make([]int, 0, len(fish))
	for _, f := range fish {
		if b := f - 1; b == -1 {
			next = append(next, 6, 8)
		} else {
			next = append(next, b)
		}
	}
	return next
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func part2(s string) uint64 {
	// First figure out the table for each number after 128 iterations
	var iter128 [9]int
	for i0 := 0; i0 < 9; i0++ {
		fish := []int{i0}
		for i := 0; i < 128; i++ {
			fish = cycle(fish)
		}
		iter128[i0] = len(fish)
	}
	// Okay now we know how long each starting number takes for 128 cycles. Now let's get for 256 cycles
	var iter256 [7]uint64
	for i0 := 0; i0 < 7; i0++ {
		fish := []int{i0}
		for i := 0; i < 128; i++ {
			fish = cycle(fish)
		}
		for _, x := range fish {
			iter256[i0] += uint64(iter128[x])
		}
	}
	// Now we know how long it takes for all of them
	var sum uint64
	for _, a := range strings.Split(s, ",") {
		sum += iter256[atoi(a)]
	}
	return sum
}
