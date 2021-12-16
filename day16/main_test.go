// Advent of Code 2021, Day 16
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		transmission string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"8A004A801A8002F478"}, 16},
		{"example2", args{"620080001611562C8802118E34"}, 12},
		{"example3", args{"C0015000016115A2E0802F182340"}, 23},
		{"example4", args{"A0016C880162017C3686B18A3D4780"}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.transmission); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		transmission string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sum", args{"C200B40A82"}, 3},
		{"product", args{"04005AC33890"}, 54},
		{"minimum", args{"880086C3E88112"}, 7},
		{"maximum", args{"CE00C43D881120"}, 9},
		{"less", args{"D8005AC2A8F0"}, 1},
		{"greater", args{"F600BC2D8F"}, 0},
		{"equal", args{"9C005AC2F8F0"}, 0},
		{"example", args{"9C0141080250320F1802104A08"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.transmission); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
