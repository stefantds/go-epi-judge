package uniform_random_number_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/uniform_random_number"
	"github.com/stefantds/go-epi-judge/stats"
)

func TestUniformRandom(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "uniform_random_number.tsv")
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := uniformRandomWrapper(tc.LowerBound, tc.UpperBound); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func uniformRandomWrapper(lowerBound int, upperBound int) error {
	return stats.RunFuncWithRetries(
		func() bool {
			return uniformRandomRunner(lowerBound, upperBound)
		},
		errors.New("the results don't match the expected distribution"),
	)
}

func uniformRandomRunner(lowerBound, upperBound int) bool {
	const nbRuns = 100000
	results := make([]int, nbRuns)

	for i := 0; i < nbRuns; i++ {
		results[i] = UniformRandom(lowerBound, upperBound)
	}

	sequence := make([]int, nbRuns)
	for i, result := range results {
		sequence[i] = result - lowerBound
	}
	return stats.CheckSequenceIsUniformlyRandom(sequence, upperBound-lowerBound+1, 0.01)
}
