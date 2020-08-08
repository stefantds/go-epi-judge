package stats

func BinomialCoefficient(n, k int) int {
	switch {
	case n < k:
		return 0
	case k == 0, k == n:
		return 1
	default:
		return (BinomialCoefficient(n-1, k-1)) + BinomialCoefficient(n-1, k)
	}
}

// Get the mth combination in lexicographical order from A (n elements) chosen
// k at a time.
func ComputeCombinationIdx(a []int, k, m int) []int {
	n := len(a)
	comb := make([]int, k)

	s, t := n, k
	x := BinomialCoefficient(n, k) - 1 - m

	for i := 0; i < k; i++ {
		s--
		for BinomialCoefficient(s, t) > x {
			s--
		}

		comb[i] = a[n-1-s]

		x -= BinomialCoefficient(s, t)
		t--
	}

	return comb
}
