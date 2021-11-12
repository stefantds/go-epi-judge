package uniform_random_number_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/uniform_random_number"
	utils "github.com/stefantds/go-epi-judge/test_utils"
	"github.com/stefantds/go-epi-judge/test_utils/stats"
)

type solutionFunc = func(int, int) int

var solutions = []solutionFunc{
	UniformRandom,
}

func TestUniformRandom(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "uniform_random_number.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		LowerBound int
		UpperBound int
		Details    string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.LowerBound,
			&tc.UpperBound,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := uniformRandomWrapper(s, tc.LowerBound, tc.UpperBound); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func uniformRandomWrapper(solution solutionFunc, lowerBound int, upperBound int) error {
	return stats.RunFuncWithRetries(
		func() bool {
			return uniformRandomRunner(solution, lowerBound, upperBound)
		},
		errors.New("the results don't match the expected distribution"),
	)
}

func uniformRandomRunner(solution solutionFunc, lowerBound, upperBound int) bool {
	const nbRuns = 100000
	results := make([]int, nbRuns)

	for i := 0; i < nbRuns; i++ {
		results[i] = solution(lowerBound, upperBound)
	}

	sequence := make([]int, nbRuns)
	for i, result := range results {
		sequence[i] = result - lowerBound
	}
	return stats.CheckSequenceIsUniformlyRandom(sequence, upperBound-lowerBound+1, 0.01)
}
