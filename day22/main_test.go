// Advent of Code 2021, Day 22
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
		{"example1", args{common.ReadStringsFromFile("testdata/example1.txt")}, 39},
		{"example2", args{common.ReadStringsFromFile("testdata/example2.txt")}, 590784},
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
		want uint64
	}{
		{"example1", args{common.ReadStringsFromFile("testdata/example1.txt")}, 39},
		{"example2", args{common.ReadStringsFromFile("testdata/example2.txt")}, 590784},
		{"example3", args{common.ReadStringsFromFile("testdata/example3.txt")}, 2758514936282235},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
