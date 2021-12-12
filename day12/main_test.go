// Advent of Code 2021, Day 12
package main

import (
	"testing"

	"github.com/ghonzo/advent2021/common"
)

func Test_part1(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{common.ReadStringsFromFile("testdata/example1.txt")}, 10},
		{"example2", args{common.ReadStringsFromFile("testdata/example2.txt")}, 19},
		{"example3", args{common.ReadStringsFromFile("testdata/example3.txt")}, 226},
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
		{"example1", args{common.ReadStringsFromFile("testdata/example1.txt")}, 36},
		{"example2", args{common.ReadStringsFromFile("testdata/example2.txt")}, 103},
		{"example3", args{common.ReadStringsFromFile("testdata/example3.txt")}, 3509},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
