// Advent of Code 2021, Day 12
package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ghonzo/advent2021/common"
)

// Day 12: Passage Pathing
// Part 1 answer: 3495
// Part 2 answer: 94849
func main() {
	fmt.Println("Advent of Code 2021, Day 12")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	start := readNodes(entries)
	return len(findPaths(path{start}, canVisit))
}

func part2(entries []string) int {
	start := readNodes(entries)
	return len(findPaths(path{start}, canVisit2))
}

// Reads the "node map" and returns the start node
func readNodes(entries []string) *node {
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
	return nodemap["start"]
}

type node struct {
	label     string
	connected []*node
}

func (n *node) isBig() bool {
	return unicode.IsUpper(rune(n.label[0]))
}

type path []*node

// Returns true if we can visit this node without violating rules (part 1)
func canVisit(p path, n *node) bool {
	if n.isBig() {
		return true
	}
	for _, cave := range p {
		if cave == n {
			return false
		}
	}
	return true
}

// Returns all the "next step" paths
func findPaths(p path, visitFunc func(path, *node) bool) []path {
	last := p[len(p)-1]
	// Recursion terminator
	if last.label == "end" {
		return []path{p}
	}
	var allPaths []path
	for _, n := range last.connected {
		if visitFunc(p, n) {
			allPaths = append(allPaths, findPaths(append(p, n), visitFunc)...)
		}
	}
	return allPaths
}

// Returns true if we can visit this node without violating rules (part 2)
func canVisit2(p path, n *node) bool {
	if n.label == "start" {
		return false
	}
	if n.isBig() {
		return true
	}
	for _, cave := range p {
		if cave == n {
			return !p.hasSmallTwice()
		}
	}
	return true
}

func (p path) hasSmallTwice() bool {
	visited := make(map[*node]bool)
	for _, n := range p {
		if !n.isBig() {
			if visited[n] {
				return true
			}
			visited[n] = true
		}
	}
	return false
}
