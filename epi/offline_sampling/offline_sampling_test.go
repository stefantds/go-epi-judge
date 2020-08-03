package offline_sampling_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stefantds/go-epi-judge/utils"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/offline_sampling"
	"github.com/stefantds/go-epi-judge/random"
)

func TestRandomSampling(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "offline_sampling.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		K       int
		A       []int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.K,
			&tc.A,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := randomSamplingWrapper(tc.K, tc.A); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func randomSamplingWrapper(k int, a []int) error {
	return random.RunFuncWithRetries(
		func() bool {
			return randomSamplingRunner(k, a)
		},
		errors.New("the results don't match the expected distribution"),
	)
}

func randomSamplingRunner(k int, a []int) bool {
	const N = 1000000

	results := make([][]int, N)

	for i := 0; i < N; i++ {
		copyA := make([]int, len(a))
		copy(copyA, a)

		RandomSampling(k, copyA)

		result := make([]int, k)
		copy(result, a[0:k])
		results[i] = result
	}

	totalPossibleOutcomes := random.BinomialCoefficient(len(a), k)

	sort.Ints(a)

	combinations := make([][]int, totalPossibleOutcomes)
	for i := 0; i < totalPossibleOutcomes; i++ {
		combinations[i] = random.ComputeCombinationIdx(a, k, i)
	}

	sort.Slice(combinations, func(i, j int) bool {
		return utils.LexIntsCompare(combinations[i], combinations[j])
	})

	sequence := make([]int, len(results))
	for i, r := range results {
		sort.Ints(r)
		sequence[i] = sort.Search(
			len(combinations),
			func(i int) bool { return !utils.LexIntsCompare(r, combinations[i]) },
		)
	}

	return random.CheckSequenceIsUniformlyRandom(sequence, totalPossibleOutcomes, 0.01)
}
