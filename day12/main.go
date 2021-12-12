// Advent of Code 2021, Day 12
package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ghonzo/advent2021/common"
)

// Day 12:
// Part 1 answer:
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2021, Day 12")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	//fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	nodemap := make(map[string]*node)
	for _, line := range entries {
		parts := strings.Split(line, "-")
		var left, right *node
		var ok bool
		if left, ok = nodemap[parts[0]]; !ok {
			left = &node{parts[0], make([]*node, 0)}
			nodemap[parts[0]] = left
		}
		if right, ok = nodemap[parts[1]]; !ok {
			right = &node{parts[1], make([]*node, 0)}
			nodemap[parts[1]] = right
		}
		left.connected = append(left.connected, right)
		right.connected = append(right.connected, left)
	}
	start := path{nodemap["start"]}
	return len(findPaths(start))
}

type node struct {
	label     string
	connected []*node
}

func (n *node) isBig() bool {
	return unicode.IsUpper(rune(n.label[0]))
}

type path []*node

func (p path) canVisit(n *node) bool {
	for _, stop := range p {
		if stop == n {
			return n.isBig()
		}
	}
	return true
}

func (p path) lastNode() *node {
	return p[len(p)-1]
}

func findPaths(p path) []path {
	last := p.lastNode()
	if last.label == "end" {
		return []path{p}
	}
	var allPaths []path
	for _, n := range last.connected {
		if p.canVisit(n) {
			allPaths = append(allPaths, findPaths(append(p, n))...)
		}
	}
	return allPaths
}
