// Advent of Code 2021, Day 4
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// The board dimension
const dim int = 5

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
	for i := 2; i < len(entries); i += dim + 1 {
		boards = append(boards, createBoard(entries[i:i+dim]))
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
	for i := 2; i < len(entries); i += dim + 1 {
		boards = append(boards, createBoard(entries[i:i+dim]))
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

type Board [dim][dim]*Space

type Space struct {
	val     int
	covered bool
}

func createBoard(lines []string) *Board {
	var b Board
	for row, line := range lines {
		for col, numStr := range strings.Fields(line) {
			val, _ := strconv.Atoi(numStr)
			b[row][col] = &Space{val: val}
		}
	}
	return &b
}

func (b *Board) Cover(v int) bool {
	// Return true if a winner
	var hcovered [dim]int
	var vcovered [dim]int
	for row, spaces := range *b {
		for col, space := range spaces {
			if space.val == v {
				space.covered = true
			}
			if space.covered {
				hcovered[row]++
				vcovered[col]++
			}
			if hcovered[row] == dim || vcovered[col] == dim {
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
