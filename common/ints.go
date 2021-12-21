package common

import "strconv"

// Abs returns the absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Atoi is just like the one in strconv, except we throw out the error
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Sgn returns 1 for a positive number, -1 for a negative number, and 0 for zero
func Sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	}
	return 0
}

// Max returns the greater of two values
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the lesser of two values
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxMin holds onto a maximum and minumum value of an arbitrary number of values.
// Create with new() and call the Accept method.
type MaxMin struct {
	Max, Min    int
	initialized bool // will be false until we get at least one value
}

// Accept a new value and update the max and min according
func (mm *MaxMin) Accept(v int) *MaxMin {
	if v > mm.Max || !mm.initialized {
		mm.Max = v
	}
	if v < mm.Min || !mm.initialized {
		mm.Min = v
	}
	mm.initialized = true
	return mm
}
