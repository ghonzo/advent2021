// Advent of Code 2021, Day 18
package main

import (
	"reflect"
	"testing"

	"github.com/ghonzo/advent2021/common"
)

func Test_parseRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{"example1", "[1,2]"},
		{"example2", "[[1,2],3]"},
		{"example3", "[9,[8,7]]"},
		{"example4", "[[1,9],[8,5]]"},
		{"example5", "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"},
		{"example6", "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"},
		{"example7", "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.s); !reflect.DeepEqual(got.String(), tt.s) {
				t.Errorf("parse() = %v, want %v", got.String(), tt.s)
			}
		})
	}
}

func Test_sf_reduce(t *testing.T) {
	tests := []struct {
		name   string
		before string
		after  string
	}{
		{"example1", "[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"example2", "[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"example3", "[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"example4", "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
		{"example5", "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := parse(tt.before)
			num.reduce()
			if got := num.String(); got != tt.after {
				t.Errorf("reduce() = %v, want %v", got, tt.after)
			}
		})
	}
}

func Test_sf_magnitude(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"example1", "[[1,2],[[3,4],5]]", 143},
		{"example2", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"example3", "[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"example4", "[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"example5", "[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"example6", "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := parse(tt.str)
			if got := num.magnitude(); got != tt.want {
				t.Errorf("sf.magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 4140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 3993},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
