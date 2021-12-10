// Advent of Code 2021, Day 10
package main

import (
	"fmt"
	"sort"

	"github.com/ghonzo/advent2021/common"
)

// Day 10:
// Part 1 answer:
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2021, Day 10")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	var sum int
lines:
	for _, line := range entries {
		var stack []rune
		for _, r := range line {
			switch r {
			case ')':
				if pop(&stack) != '(' {
					sum += 3
					continue lines
				}
			case ']':
				if pop(&stack) != '[' {
					sum += 57
					continue lines
				}
			case '}':
				if pop(&stack) != '{' {
					sum += 1197
					continue lines
				}
			case '>':
				if pop(&stack) != '<' {
					sum += 25137
					continue lines
				}
			default:
				stack = append(stack, r)
			}
		}
	}
	return sum
}

func part2(entries []string) uint64 {
	var scores []uint64
lines:
	for _, line := range entries {
		var stack []rune
		for _, r := range line {
			switch r {
			case ')':
				if pop(&stack) != '(' {
					continue lines
				}
			case ']':
				if pop(&stack) != '[' {
					continue lines
				}
			case '}':
				if pop(&stack) != '{' {
					continue lines
				}
			case '>':
				if pop(&stack) != '<' {
					continue lines
				}
			default:
				stack = append(stack, r)
			}
		}
		var sum uint64
		for i := len(stack) - 1; i >= 0; i-- {
			sum *= 5
			switch stack[i] {
			case '(':
				sum += 1
			case '[':
				sum += 2
			case '{':
				sum += 3
			case '<':
				sum += 4
			default:
				panic("noes")
			}
		}
		scores = append(scores, sum)
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	//sort.Ints(scores)
	return scores[len(scores)/2]
}

func pop(s *[]rune) rune {
	n := len(*s) - 1
	v := (*s)[n]
	*s = (*s)[:n]
	return v
}
