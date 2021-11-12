package random_subset_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/random_subset"
	utils "github.com/stefantds/go-epi-judge/test_utils"
	"github.com/stefantds/go-epi-judge/test_utils/stats"
)

type solutionFunc = func(int, int) []int

var solutions = []solutionFunc{
	RandomSubset,
}

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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := randomSubsetWrapper(s, tc.N, tc.K); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func randomSubsetWrapper(solution solutionFunc, n int, k int) error {
	return stats.RunFuncWithRetries(
		func() bool {
			return randomSubsetRunner(solution, n, k)
		},
		errors.New("the results don't match the expected distribution"),
	)
}

func randomSubsetRunner(solution solutionFunc, n int, k int) bool {
	const nbRuns = 1000000
	results := make([][]int, nbRuns)

	for i := 0; i < nbRuns; i++ {
		results[i] = solution(n, k)
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
		pos := sort.Search(
			len(combinations),
			func(i int) bool { return !utils.LexIntsCompare(combinations[i], r) },
		)
		if pos < len(combinations) && reflect.DeepEqual(combinations[pos], r) {
			sequence[i] = pos
		} else {
			panic("result not in known combinations")
		}
	}

	return stats.CheckSequenceIsUniformlyRandom(sequence, totalPossibleOutcomes, 0.01)
}
