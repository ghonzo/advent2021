// Advent of Code 2021, Day 22
package main

import (
	"fmt"
	"regexp"

	"github.com/ghonzo/advent2021/common"
)

// Day 22: Reactor Reboot
// Part 1 answer: 611378
// Part 2 answer: 1214313344725528
func main() {
	fmt.Println("Advent of Code 2021, Day 22")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

type cube struct {
	on     bool
	bounds [3][2]int // x, y, z
}

func part1(entries []string) int {
	var space [101][101][101]bool // [x][y][z]
	for _, rs := range parseSteps(entries) {
		for x := common.Max(rs.bounds[0][0], -50); x <= common.Min(rs.bounds[0][1], 50); x++ {
			for y := common.Max(rs.bounds[1][0], -50); y <= common.Min(rs.bounds[1][1], 50); y++ {
				for z := common.Max(rs.bounds[2][0], -50); z <= common.Min(rs.bounds[2][1], 50); z++ {
					space[x+50][y+50][z+50] = rs.on
				}
			}
		}
	}
	var count int
	for x := 0; x < 101; x++ {
		for y := 0; y < 101; y++ {
			for z := 0; z < 101; z++ {
				if space[x][y][z] {
					count++
				}
			}
		}
	}
	return count
}

func parseSteps(entries []string) []cube {
	r := regexp.MustCompile(`([^ ]+) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)
	var steps []cube
	for _, line := range entries {
		rs := cube{}
		groups := r.FindStringSubmatch(line)
		rs.on = (groups[1] == "on")
		for dim := 0; dim < 3; dim++ {
			rs.bounds[dim][0] = common.Atoi(groups[2+dim*2])
			rs.bounds[dim][1] = common.Atoi(groups[3+dim*2])
		}
		steps = append(steps, rs)
	}
	return steps
}

func part2(entries []string) uint64 {
	steps := parseSteps(entries)
	var cubes []cube
	for _, s := range steps {
		var cubesToAdd []cube
		// Check to see if there's an intersection with any previous cubes
		for _, c := range cubes {
			if ic, ok := intersection(s, c); ok {
				// Yes there was, so we need to treat this like a new cube that is the "inverse" of what we intersected with
				ic.on = !c.on
				cubesToAdd = append(cubesToAdd, ic)
			}
		}
		// Okay, took care of all the intersections. Add all the intersections cubes to our list
		cubes = append(cubes, cubesToAdd...)
		// And if this step was an "on" step, add that cube to the list. Leave it off it was an "off" cube
		if s.on {
			cubes = append(cubes, s)
		}
	}
	// Okay, now we have all of our cubes, positive and negative
	var sum uint64
	for _, c := range cubes {
		v := c.volume()
		if c.on {
			sum += v
		} else {
			sum -= v
		}
	}
	return sum
}

func intersection(c1, c2 cube) (cube, bool) {
	var ic [3][2]int
	for axis := 0; axis < 3; axis++ {
		ic[axis][0] = common.Max(c1.bounds[axis][0], c2.bounds[axis][0])
		ic[axis][1] = common.Min(c1.bounds[axis][1], c2.bounds[axis][1])
	}
	return cube{bounds: ic}, ic[0][0] <= ic[0][1] && ic[1][0] <= ic[1][1] && ic[2][0] <= ic[2][1]
}

func (c cube) volume() uint64 {
	// NOTE THE OFF-BY-ONE POTENTIAL GOTCHA!
	return uint64((c.bounds[0][1] - c.bounds[0][0] + 1) * (c.bounds[1][1] - c.bounds[1][0] + 1) * (c.bounds[2][1] - c.bounds[2][0] + 1))
}
