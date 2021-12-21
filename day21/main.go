// Advent of Code 2021, Day 21
package main

import (
	"fmt"
)

// Day 21: Dirac Dice
// Part 1 answer: 752745
// Part 2 answer: 309196008717909
func main() {
	fmt.Println("Advent of Code 2021, Day 21")
	// 6 and 3 come from the input
	fmt.Printf("Part 1: %d \n", part1(6, 3))
	fmt.Printf("Part 2: %d \n", part2(6, 3))
}

func part1(p1, p2 int) int {
	pos := [2]int{p1, p2}
	die := 1
	var score [2]int
	player := 0
	rollCount := 0
	for {
		roll := die
		if die = die + 1; die > 100 {
			die = 1
		}
		roll += die
		if die = die + 1; die > 100 {
			die = 1
		}
		roll += die
		if die = die + 1; die > 100 {
			die = 1
		}
		rollCount += 3
		newBoardSpace := (pos[player]+roll-1)%10 + 1
		pos[player] = newBoardSpace
		score[player] += newBoardSpace
		if score[player] >= 1000 {
			return score[1-player] * rollCount
		}
		player = 1 - player
	}
}

type state struct {
	pos        [2]int
	score      [2]int
	nextPlayer int
}

type wins [2]uint64

func part2(p1, p2 int) uint64 {
	universes := make(map[state]wins)
	initialState := state{pos: [2]int{p1, p2}}
	w := playGame(initialState, universes)
	if w[0] > w[1] {
		return w[0]
	}
	return w[1]
}

func playGame(s state, universes map[state]wins) wins {
	// Have we seen this state before?
	if w, ok := universes[s]; ok {
		return w
	}
	var w wins
	player := s.nextPlayer
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				roll := d1 + d2 + d3
				newState := s
				newState.pos[player] = (s.pos[player]-1+roll)%10 + 1
				newState.score[player] += newState.pos[player]
				if newState.score[player] >= 21 {
					w[player]++
				} else {
					newState.nextPlayer = 1 - player
					subWins := playGame(newState, universes)
					w[0] += subWins[0]
					w[1] += subWins[1]
				}
			}
		}
	}
	// remember it
	universes[s] = w
	return w
}
