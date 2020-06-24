package epi_test

import (
	"errors"
	"fmt"
	"math"
	"os"
	"testing"

	. "github.com/stefantds/go-epi-judge/epi"
	"github.com/stefantds/go-epi-judge/random"

	csv "github.com/stefantds/csvdecoder"
)

func TestNonuniformRandomNumberGeneration(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "nonuniform_random_number.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Values        []int
		Probabilities []float64
		Details       string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Values,
			&tc.Probabilities,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := nonuniformRandomNumberGenerationWrapper(tc.Values, tc.Probabilities); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func nonuniformRandomNumberGenerationWrapper(values []int, probabilities []float64) error {
	return random.RunFuncWithRetries(
		func() bool {
			return nonuniformRandomNumberGenerationRunner(values, probabilities)
		},
		errors.New("the generation doesn't match the expected distribution"),
	)
}

func nonuniformRandomNumberGenerationRunner(values []int, probabilities []float64) bool {
	const N = 1000000

	results := make([]int, N)
	for i := 0; i < N; i++ {
		results[i] = NonuniformRandomNumberGeneration(values, probabilities)
	}

	counts := make(map[int]int, len(values))
	for _, r := range results {
		counts[r] += 1
	}

	for i, v := range values {
		p := probabilities[i]
		if N*p < 50 || N*(1.0-p) < 50 {
			continue
		}

		sigma := math.Sqrt(N * p * (1.0 - p))

		if math.Abs(float64(counts[v])-(p*N)) > 5*sigma {
			return false
		}
	}
	return true
}
