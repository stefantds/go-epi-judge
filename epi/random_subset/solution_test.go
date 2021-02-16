package random_subset_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/random_subset"
	utils "github.com/stefantds/go-epi-judge/test_utils"
	"github.com/stefantds/go-epi-judge/test_utils/stats"
)

func TestRandomSubset(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "random_subset.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		N       int
		K       int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.N,
			&tc.K,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			if err := randomSubsetWrapper(tc.N, tc.K); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func randomSubsetWrapper(n int, k int) error {
	return stats.RunFuncWithRetries(
		func() bool {
			return randomSubsetRunner(n, k)
		},
		errors.New("the results don't match the expected distribution"),
	)
}

func randomSubsetRunner(n int, k int) bool {
	const nbRuns = 1000000
	results := make([][]int, nbRuns)

	for i := 0; i < nbRuns; i++ {
		results[i] = RandomSubset(n, k)
	}

	totalPossibleOutcomes := stats.BinomialCoefficient(n, k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}

	combinations := make([][]int, totalPossibleOutcomes)
	for i := 0; i < totalPossibleOutcomes; i++ {
		combinations[i] = stats.ComputeCombinationIdx(a, k, i)
	}

	sort.Slice(combinations, func(i, j int) bool {
		return utils.LexIntsCompare(combinations[i], combinations[j])
	})

	sequence := make([]int, nbRuns)
	for i, r := range results {
		sort.Ints(r)
		sequence[i] = sort.Search(
			len(combinations),
			func(i int) bool { return !utils.LexIntsCompare(r, combinations[i]) },
		)
	}
	return stats.CheckSequenceIsUniformlyRandom(sequence, totalPossibleOutcomes, 0.01)
}
