// Advent of Code 2021, Day 1
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		entries []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, 7},
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
		entries []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
