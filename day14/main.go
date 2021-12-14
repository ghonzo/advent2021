// Advent of Code 2021, Day 14
package main

import (
	"fmt"

	"github.com/ghonzo/advent2021/common"
)

// Day 14: Extended Polymerization
// Part 1 answer: 2321
// Part 2 answer: 2399822193707
func main() {
	fmt.Println("Advent of Code 2021, Day 14")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d \n", part1(entries))
	fmt.Printf("Part 2: %d \n", part2(entries))
}

func part1(entries []string) int {
	template := []byte(entries[0])
	rules := make(map[string]byte)
	for _, s := range entries[2:] {
		rules[s[0:2]] = byte(s[6])
	}
	for n := 0; n < 10; n++ {
		newTemplate := []byte{template[0]}
		for i := 0; i < len(template)-1; i++ {
			newTemplate = append(newTemplate, rules[string(template[i:i+2])], template[i+1])
		}
		template = newTemplate
	}
	distMap := distribution(template)
	return int(Range(distMap))
}

func distribution(s []byte) map[byte]uint64 {
	countMap := make(map[byte]uint64)
	for _, b := range s {
		countMap[b]++
	}
	return countMap
}

// Range is the difference between the maximum and the minimum of the values
func Range(m map[byte]uint64) uint64 {
	min := ^uint64(0)
	max := uint64(0)
	for _, v := range m {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func part2(entries []string) uint64 {
	// Find the distribution map after 20 iterations for each two char rule
	rules := make(map[string]byte)
	for _, s := range entries[2:] {
		rules[s[0:2]] = byte(s[6])
	}
	// For each pair, the char distribution after 20 iters
	pairDistMap20 := make(map[string]map[byte]uint64)
	// What each rule looks like after 20 iterations
	templates20 := make(map[string][]byte)
	for pair := range rules {
		template := []byte(pair)
		for n := 0; n < 20; n++ {
			newTemplate := []byte{template[0]}
			for i := 0; i < len(template)-1; i++ {
				newTemplate = append(newTemplate, rules[string(template[i:i+2])], template[i+1])
			}
			template = newTemplate
		}
		templates20[pair] = template
		// Now figure out the char dist
		pairDistMap20[pair] = distribution(template)
	}
	// Now iterate over each template20 and aggregate the distMaps
	pairDistMap40 := make(map[string]map[byte]uint64)
	for pair, template := range templates20 {
		distMap40 := make(map[byte]uint64)
		for i := 0; i < len(template)-1; i++ {
			addMaps(distMap40, pairDistMap20[(string(template[i:i+2]))])
		}
		pairDistMap40[pair] = distMap40
	}
	// Finally iterate over the initial state
	initialState := entries[0]
	uberMap := distribution([]byte(initialState))
	for i := 0; i < len(initialState)-1; i++ {
		addMaps(uberMap, pairDistMap40[(string(initialState[i:i+2]))])
		// Otherwise we count iter 20 twice
		subMaps(uberMap, pairDistMap20[(string(initialState[i:i+2]))])
	}
	return Range(uberMap)
}

func addMaps(m1, m2 map[byte]uint64) {
	for k, v := range m2 {
		m1[k] += v
	}
}

func subMaps(m1, m2 map[byte]uint64) {
	for k, v := range m2 {
		m1[k] -= v
	}
}
