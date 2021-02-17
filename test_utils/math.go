package test_utils

import "math"

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// Max returns the greater value of two integers.
// Nothing fancy, just a convenience function.
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller value of two integers.
// Nothing fancy, just a convenience function.
func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// EqualFloat returns true if the two float values are "equal" within
// the tolerance value
func EqualFloat(a, b float64) bool {
	tolerance := 0.00001
	return math.Abs(a-b) < tolerance
}
