// Advent of Code 2021, Day 18
package main

import (
	"fmt"
	"strconv"

	"github.com/ghonzo/advent2021/common"
)

// Day 18: Snailfish
// Part 1 answer: 4391
// Part 2 answer: 4626
func main() {
	fmt.Println("Advent of Code 2021, Day 18")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	sum := parse(entries[0])
	for _, line := range entries[1:] {
		sum = addAndReduce(sum, parse(line))
	}
	return sum.magnitude()
}

func part2(entries []string) int {
	maxMag := 0
	for i, a := range entries[:len(entries)-1] {
		for _, b := range entries[i+1:] {
			mag := addAndReduce(parse(a), parse(b)).magnitude()
			if mag > maxMag {
				maxMag = mag
			}
			// Swap the order
			mag = addAndReduce(parse(b), parse(a)).magnitude()
			if mag > maxMag {
				maxMag = mag
			}
		}
	}
	return maxMag
}

func parse(s string) *node {
	index := 0
	return parseNum(s, &index)
}

func parseNum(s string, i *int) *node {
	if s[*i] == '[' {
		*i++
		num := new(node)
		num.left = parseNum(s, i)
		// Next char should be ','
		if s[*i] != ',' {
			panic("expected a ,")
		}
		*i++
		num.right = parseNum(s, i)
		// Next char should be a ']'
		if s[*i] != ']' {
			panic("expected a ]")
		}
		*i++
		return num
	}
	// When we parse it's always a single number
	v := s[*i] - '0'
	*i++
	return &node{val: int(v)}
}

// Rather than mess around with two kinds of nodes (regular or pair), we're just going lazy and combining them
// into a single stuct. If left and right are nil, then it's a "regular" node with the given val. Otherwise, it's
// a pair and val is unused.
type node struct {
	val         int
	left, right *node
}

func (num *node) String() string {
	if num.isRegular() {
		return strconv.Itoa(num.val)
	}
	return fmt.Sprintf("[%v,%v]", num.left, num.right)
}

func (num *node) isRegular() bool {
	return num.left == nil && num.right == nil
}

func addAndReduce(n1, n2 *node) *node {
	newParent := &node{left: n1, right: n2}
	newParent.reduce()
	return newParent
}

func (num *node) reduce() {
	// Keep going until these both return false
	for explode(num) || split(num) {
	}
}

type nodeVisitor interface {
	// Return true if we should keep traversing
	accept(n *node, depth int) bool
}

// Returns true if we should continue to traverse, or at the top level if the tree was not modified
func traverse(n *node, depth int, visitor nodeVisitor) bool {
	c := visitor.accept(n, depth)
	if c && !n.isRegular() {
		if c = traverse(n.left, depth+1, visitor); c {
			c = traverse(n.right, depth+1, visitor)
		}
	}
	return c
}

type explodeVisitor struct {
	lastRegular *node
	exploded    bool
	numToAdd    int
}

func (ev *explodeVisitor) accept(n *node, depth int) bool {
	if n.isRegular() {
		// If we've already exploded, then add numToAdd to this node
		if ev.exploded {
			n.val += ev.numToAdd
			// We're done
			return false
		}
		// Save this node for future explosions
		ev.lastRegular = n
		return true
	}
	// If depth is 4 and we haven't exploded, then BOOM
	if depth == 4 && !ev.exploded {
		ev.exploded = true
		// Add the left value to the last regular node we've seen
		if ev.lastRegular != nil {
			ev.lastRegular.val += n.left.val
		}
		// And save the right value for the next regular node ... unless it's zero
		ev.numToAdd = n.right.val
		// Make this a regular node of value zero (which is what val is already)
		n.left = nil
		n.right = nil
		return ev.numToAdd > 0 // If it's zero we can just stop
	}
	return true
}

func explode(n *node) bool {
	ev := new(explodeVisitor)
	traverse(n, 0, ev)
	return ev.exploded
}

type splitVisitor struct{}

func (sv splitVisitor) accept(n *node, depth int) bool {
	if n.isRegular() {
		if v := n.val; v >= 10 {
			// Turn it into a pair and stop
			n.val = 0
			n.left = &node{val: v / 2}
			n.right = &node{val: (v + 1) / 2}
			return false
		}
	}
	return true
}

func split(n *node) bool {
	return !traverse(n, 0, splitVisitor{})
}

func (num *node) magnitude() int {
	if num.isRegular() {
		return num.val
	}
	return num.left.magnitude()*3 + num.right.magnitude()*2
}
