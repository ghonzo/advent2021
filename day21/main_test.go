// Advent of Code 2021, Day 21
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		p1 int
		p2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{4, 8}, 739785},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		p1 int
		p2 int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"example", args{4, 8}, 444356092776315},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
