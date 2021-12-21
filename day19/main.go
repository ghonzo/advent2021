// Advent of Code 2021, Day 19
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2021/common"
)

// Day 19: Beacon Scanner
// Part 1 answer: 390
// Part 2 answer: 13327
func main() {
	fmt.Println("Advent of Code 2021, Day 19")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

type point3d [3]int

type scanner struct {
	beacons    []point3d
	location   point3d
	normalized bool
}

type segment [2]point3d

func part1(entries []string) int {
	scanners := calculateScanners(entries)
	// Now collect all the unique beacon locations
	beaconLocations := make(map[point3d]int)
	for _, s := range scanners {
		for _, p := range s.beacons {
			beaconLocations[p.add(s.location)]++
		}
	}
	return len(beaconLocations)
}

func part2(entries []string) int {
	scanners := calculateScanners(entries)
	return findMaxDistance(scanners)
}

func calculateScanners(entries []string) []scanner {
	scanners := parseScanners(entries)
	scanners[0].normalized = true
	for doMoreWork(scanners) {
		normalizeScanners(scanners)
	}
	return scanners
}

func parseScanners(entries []string) []scanner {
	var scanners []scanner
	var curScanner scanner
	skip := true
	for _, line := range entries {
		if skip {
			skip = false
		} else if len(line) == 0 {
			scanners = append(scanners, curScanner)
			curScanner = scanner{}
			skip = true
		} else {
			n := strings.Split(line, ",")
			curScanner.beacons = append(curScanner.beacons, point3d{common.Atoi(n[0]), common.Atoi(n[1]), common.Atoi(n[2])})
		}
	}
	return append(scanners, curScanner)
}

// Returns true if there's at least one scanner that hasn't been normalized yet
func doMoreWork(scanners []scanner) bool {
	for _, s := range scanners {
		if !s.normalized {
			return true
		}
	}
	return false
}

// each value represents a "dimension mapping" based on a three-element vector, so [1,2,3] is
// the identity translation (like [x,y,z]) while [2,1,-3] is like [x,y,z] -> [y,x,-z]
type translationVector [3]int

func normalizeScanners(scanners []scanner) {
	for i := 0; i < len(scanners); i++ {
		s := scanners[i]
		if !s.normalized {
			continue
		}
		// Find other scanners that overlap with this one that haven't been mapped yet
		segments := findAllSegments(s)
		for j := 0; j < len(scanners); j++ {
			s2 := scanners[j]
			if i == j || s2.normalized {
				continue
			}
			segments2 := findAllSegments(s2)
			overlaps := 0
			// Stores a count of the most popular offset
			offsetCount := make(map[point3d]int)
			// Now for each offset, store the associated translations. Might get overwritten a lot, that's fine
			offsetToTranslationMap := make(map[point3d]translationVector)
			for k := range segments {
				if _, ok := segments2[k]; ok {
					if translation, offset, ok := findTranslationAndOffset(segments[k], segments2[k]); ok {
						overlaps++
						offsetCount[offset]++
						offsetToTranslationMap[offset] = translation
					}
				}
			}
			// If there are 66 shared connections between nodes, that represents 12 shared nodes via n(n-1)/2
			if overlaps >= 66 {
				// Find the most popular offset
				var offset point3d
				max := 0
				for k, v := range offsetCount {
					if v > max {
						offset = k
						max = v
					}
				}
				// Now find the translation for this offset point
				translation := offsetToTranslationMap[offset]
				scanners[j].location = offset.add(s.location)
				// And fix all the scanner's beacon locations
				for bi, p := range s2.beacons {
					scanners[j].beacons[bi] = p.translated(translation)
				}
				scanners[j].normalized = true
				return
			}
		}
	}
}

func findAllSegments(s scanner) map[int]segment {
	distanceMap := make(map[int]segment)
	for i, p1 := range s.beacons[:len(s.beacons)-1] {
		for _, p2 := range s.beacons[i+1:] {
			seg := segment{p1, p2}
			distanceMap[seg.distance()] = seg
		}
	}
	return distanceMap
}

// This is actually squared distance
func (s segment) distance() int {
	d := s.diffs()
	return sqr(d[0]) + sqr(d[1]) + sqr(d[2])
}

func (s segment) diffs() [3]int {
	return [3]int{s[0][0] - s[1][0], s[0][1] - s[1][1], s[0][2] - s[1][2]}
}

func findTranslationAndOffset(s1, s2 segment) (translation translationVector, offset point3d, ok bool) {
	diffs1 := s1.diffs()
	diffs2 := s2.diffs()
outer:
	for i, d1 := range diffs1 {
		for j, d2 := range diffs2 {
			if d1 == d2 {
				translation[i] = j + 1
				continue outer
			}
			if d1 == -d2 {
				translation[i] = -(j + 1)
				continue outer
			}
		}
	}
	if translation[0] != 0 && translation[1] != 0 && translation[2] != 0 {
		ok = true
		offset = s1[0].sub(s2[0].translated(translation))
	}
	return
}

func findMaxDistance(scanners []scanner) int {
	mm := new(common.MaxMin)
	for i := 0; i < len(scanners)-1; i++ {
		p1 := scanners[i].location
		for j := i; j < len(scanners); j++ {
			p2 := scanners[j].location
			mm.Accept(p1.sub(p2).manhattanDistance())
		}
	}
	return mm.Max
}

func (p1 point3d) sub(p2 point3d) point3d {
	return point3d{p1[0] - p2[0], p1[1] - p2[1], p1[2] - p2[2]}
}

func (p1 point3d) add(p2 point3d) point3d {
	return point3d{p1[0] + p2[0], p1[1] + p2[1], p1[2] + p2[2]}
}

func (p point3d) manhattanDistance() int {
	return common.Abs(p[0]) + common.Abs(p[1]) + common.Abs(p[2])
}

// This is kind of magicky
func (p point3d) translated(t translationVector) point3d {
	var tp point3d
	for i := 0; i < 3; i++ {
		tp[i] = p[common.Abs(t[i])-1]
		if t[i] < 0 {
			tp[i] *= -1
		}
	}
	return tp
}

func sqr(n int) int {
	return n * n
}
