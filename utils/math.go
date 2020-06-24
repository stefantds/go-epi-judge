package utils

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
