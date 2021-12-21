// Advent of Code 2021, Day 17
package main

import (
	"fmt"
	"math"

	"github.com/ghonzo/advent2021/common"
)

// Day 17: Trick Shot
// Part 1 answer: 4005
// Part 2 answer: 2953
func main() {
	fmt.Println("Advent of Code 2021, Day 17")
	// This is from input
	r := rect{240, 292, -90, -57}
	fmt.Printf("Part 1: %d \n", part1(r))
	fmt.Printf("Part 2: %d \n", part2(r))
}

type rect struct {
	x0, x1, y0, y1 int
}

func part1(target rect) int {
	vxMin, vxMax, vyMin, vyMax := vRange(target)
	mmY := new(common.MaxMin)
	for vx := vxMin; vx <= vxMax; vx++ {
		for vy := vyMin; vy <= vyMax; vy++ {
			mmY.Accept(fire(vx, vy, target))
		}
	}
	return mmY.Max
}

func part2(target rect) int {
	vxMin, vxMax, vyMin, vyMax := vRange(target)
	var hits int
	for vx := vxMin; vx <= vxMax; vx++ {
		for vy := vyMin; vy <= vyMax; vy++ {
			my := fire(vx, vy, target)
			if my != missed {
				hits++
			}

		}
	}
	return hits
}

// Calculate the minimum and maximum int velocities to have a chance to hit the range
func vRange(target rect) (vxMin, vxMax, vyMin, vyMax int) {
	// Min x: d = v0*t + 0.5*a*t, where d = x0, a=-1, and t = v0. Solve for v0
	vxMin = int(math.Sqrt(2 * float64(target.x0)))
	// Max x: cannot be greater than x1
	vxMax = target.x1
	// Min y: cannot be less than y0
	vyMin = target.y0
	// Max y: Um, let's say double the abs value of y0. This is a hand wave.
	vyMax = -target.y0 * 2
	return
}

const missed = math.MinInt

// Returns the maximum height reached, or "missed" if it didn't hit the range
func fire(vx, vy int, target rect) int {
	mmY := new(common.MaxMin)
	var x, y int
	for x <= target.x1 && y >= target.y0 {
		x += vx
		y += vy
		mmY.Accept(y)
		if x >= target.x0 && x <= target.x1 && y >= target.y0 && y <= target.y1 {
			// Hit!
			return mmY.Max
		}
		vx -= common.Sgn(vx)
		vy--
	}
	return missed
}
