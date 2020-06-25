package random_permutation_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/random_permutation"
	"github.com/stefantds/go-epi-judge/random"
)

func TestComputeRandomPermutation(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "random_permutation.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		N       int
		Details string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.N,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := computeRandomPermutationWrapper(tc.N); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func computeRandomPermutationWrapper(n int) error {
	return random.RunFuncWithRetries(
		func() bool {
			return computeRandomPermutationRunner(n)
		},
		errors.New("the results don't match the expected distribution"),
	)
}

func computeRandomPermutationRunner(n int) bool {
	const nbRuns = 1000000

	results := make([][]int, nbRuns)
	for i := 0; i < nbRuns; i++ {
		results[i] = ComputeRandomPermutation(n)
	}

	sequence := make([]int, nbRuns)
	for i, r := range results {
		sequence[i] = permutationIndex(r)
	}
	return random.CheckSequenceIsUniformlyRandom(sequence, factorial(n), 0.01)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func permutationIndex(perm []int) int {
	idx := 0
	n := len(perm)
	for i := 0; i < len(perm); i++ {
		a := perm[i]
		idx += a * factorial(n-1)
		for j := i + 1; j < len(perm); j++ {
			if perm[j] > a {
				perm[j] = perm[j] - 1
			}
		}
		n--
	}
	return idx
}
