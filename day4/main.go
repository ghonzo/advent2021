// Advent of Code 2021, Day 4
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 4: Giant Squid
// Part 1 answer: 51034
// Part 2 answer: 5434
func main() {
	fmt.Println("Advent of Code 2021, Day 4")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	// First line is the numbers drawn
	numbers := strings.Split(entries[0], ",")
	boards := make([]*Board, 0)
	// Now make boards
	for i := 2; i < len(entries); i += 6 {
		boards = append(boards, createBoard(entries[i:i+5]))
	}
	// Now draw the numbers
	for _, num := range numbers {
		v, _ := strconv.Atoi(num)
		for _, b := range boards {
			if b.Cover(v) {
				return b.SumUnmarked() * v
			}
		}
	}
	panic("Didn't find a winner")
}

func part2(entries []string) int {
	// First line is the numbers drawn
	numbers := strings.Split(entries[0], ",")
	boards := make([]*Board, 0)
	// Now make boards
	for i := 2; i < len(entries); i += 6 {
		boards = append(boards, createBoard(entries[i:i+5]))
	}
	// Now draw the numbers
	winners := make(map[int]bool)
	for _, num := range numbers {
		v, _ := strconv.Atoi(num)
		for num, b := range boards {
			if b.Cover(v) {
				winners[num] = true
				if len(winners) == len(boards) {
					return b.SumUnmarked() * v
				}
			}
		}
	}
	panic("Didn't find a last winner")
}

type Board [][]*Space

type Space struct {
	val     int
	covered bool
}

func createBoard(lines []string) *Board {
	var b Board
	for _, line := range lines {
		spaces := make([]*Space, 0, 5)
		for _, numStr := range strings.Fields(line) {
			val, _ := strconv.Atoi(numStr)
			spaces = append(spaces, &Space{val: val})
		}
		b = append(b, spaces)
	}
	return &b
}

func (b *Board) Cover(v int) bool {
	// Return true if a winner
	var hcovered [5]int
	var vcovered [5]int
	for row, spaces := range *b {
		for col, space := range spaces {
			if space.val == v {
				space.covered = true
			}
			if space.covered {
				hcovered[row]++
				vcovered[col]++
			}
			if hcovered[row] == 5 || vcovered[col] == 5 {
				return true
			}
		}
	}
	return false
}

func (b *Board) SumUnmarked() int {
	var sum int
	for _, spaces := range *b {
		for _, space := range spaces {
			if !space.covered {
				sum += space.val
			}
		}
	}
	return sum
}
