// Advent of Code 2021, Day 10
package main

import (
	"testing"

	"github.com/ghonzo/advent2021/common"
)

func Test_part2(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
