package utils

// LexicographicalArrayComparator compares two int arrays using lexicographical rules.
// Returns true if a < b
func LexicographicalArrayComparator(a, b []int) bool {
	for i := 0; i < Min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}

	// one array is included in the other
	return len(a) < len(b)
}
