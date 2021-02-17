package stats

import "math"

// seq is a sequence of integers, which should be in the range [0,n-1]. We
// assume n << len(seq).
func CheckSequenceIsUniformlyRandom(seq []int, n int, falseNegativeTolerance float64) bool {
	return CheckFrequencies(seq, n, falseNegativeTolerance) &&
		CheckPairsFrequencies(seq, n, falseNegativeTolerance) &&
		CheckTriplesFrequencies(seq, n, falseNegativeTolerance) &&
		CheckBirthdaySpacings(seq, n)
}

func CheckFrequencies(seq []int, n int, falseNegativeTolerance float64) bool {
	lenSeq := float64(len(seq))
	avg := lenSeq / float64(n)
	kIndiv := float64(ComputeDeviationMultiplier(falseNegativeTolerance, n))
	p := 1.0 / float64(n)
	sigmaIndiv := math.Sqrt(lenSeq * p * (1.0 - p))
	kSigmaIndiv := kIndiv * sigmaIndiv

	// To make our testing meaningful "sufficiently large", we need to have enough testing data.
	if lenSeq*p < 50 || lenSeq*(1-p) < 50 {
		return true // Sample size is too small so we cannot use normal approximation
	}

	indivFreqs := make(map[int]int)
	for _, a := range seq {
		indivFreqs[a] += 1
	}

	// Check that there are roughly len(seq)/n occurrences of key. By roughly
	// we mean the difference is less than k_sigma.
	for _, freq := range indivFreqs {
		if !(math.Abs(avg-float64(freq)) <= kSigmaIndiv) {
			return false
		}
	}

	return true
}

func CheckPairsFrequencies(seq []int, n int, falseNegativeTolerance float64) bool {
	seqPairs := make([]int, len(seq))
	for i := 1; i < len(seq); i++ {
		seqPairs[i] = seq[i-1]*n + seq[i]
	}

	return CheckFrequencies(seqPairs, n*n, falseNegativeTolerance)
}

func CheckTriplesFrequencies(seq []int, n int, falseNegativeTolerance float64) bool {
	seqTriplets := make([]int, len(seq))
	for i := 2; i < len(seq); i++ {
		seqTriplets[i] = seq[i-2]*n*n + seq[i-1]*n + seq[i]
	}

	return CheckFrequencies(seqTriplets, n*n*n, falseNegativeTolerance)
}

func CheckBirthdaySpacings(seq []int, n int) bool {
	const (
		minNumberSubarrays = 1000
		countTolerance     = 0.4
	)

	expectedAvgRepetitionLength := int(math.Ceil(math.Sqrt(math.Log(2.0) * 2.0 * float64(n))))
	numberOfSubarrays := float64(len(seq) - expectedAvgRepetitionLength + 1)

	if numberOfSubarrays < minNumberSubarrays {
		return true // Not enough subarrays for birthday spacing check
	}

	numberOfSubarraysWithRepetitions := float64(0)

	for i := 0; i < len(seq)-expectedAvgRepetitionLength; i++ {
		seqWindow := make(map[int]bool)
		for _, s := range seq[i : i+expectedAvgRepetitionLength] {
			seqWindow[s] = true
		}
		if len(seqWindow) < expectedAvgRepetitionLength {
			numberOfSubarraysWithRepetitions += 1
		}
	}

	return countTolerance*numberOfSubarrays <= numberOfSubarraysWithRepetitions
}

func ComputeDeviationMultiplier(allowedFalseNegative float64, numRvs int) int {
	individualRvError := allowedFalseNegative / float64(numRvs)
	errorBounds := []float64{
		1 - 0.682689492137086,
		1 - 0.954499736103642,
		1 - 0.997300203936740,
		1 - 0.999936657516334,
		1 - 0.999999426696856,
		1 - 0.999999998026825,
		1 - 0.999999999997440,
	}

	for i, errBound := range errorBounds {
		if errBound <= individualRvError {
			return i + 1
		}
	}
	return len(errorBounds) + 1
}
