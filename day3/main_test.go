// Advent of Code 2021, Day 3
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		entries []string
		bits    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", args{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}, 5}, 198},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries, tt.args.bits); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		entries []string
		bits    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"sample", args{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}, 5}, 230},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries, tt.args.bits); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
