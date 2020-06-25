package utils

// LexIntsCompare compares two int arrays using lexicographical order.
// Returns true if a < b in lexicographical order.
func LexIntsCompare(a, b []int) bool {
	for i := 0; i < Min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}

	// one array is included in the other
	return len(a) < len(b)
}

// LexStringsCompare compares two string arrays using lexicographical order.
// Returns true if a < b in lexicographical order.
func LexStringsCompare(a, b []string) bool {
	for i := 0; i < Min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}

	// one array is included in the other
	return len(a) < len(b)
}
