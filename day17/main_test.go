// Advent of Code 2021, Day 17
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		target rect
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{rect{20, 30, -10, -5}}, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.target); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		target rect
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{rect{20, 30, -10, -5}}, 112},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.target); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
