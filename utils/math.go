package utils

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
