package epi_test

import (
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestFindLongestIncreasingSubarray(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "longest_increasing_subarray.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A              []int
		ExpectedLength int
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.A,
			&tc.ExpectedLength,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := findLongestIncreasingSubarrayWrapper(tc.A, tc.ExpectedLength); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findLongestIncreasingSubarrayWrapper(a []int, expectedLength int) error {
	result := FindLongestIncreasingSubarray(a)

	switch {
	case result.Start < 0 || result.Start >= len(a):
		return fmt.Errorf("invalid start index %d", result.Start)
	case result.End < 0 || result.End >= len(a):
		return fmt.Errorf("invalid end index %d", result.End)
	case result.End < result.Start:
		return fmt.Errorf("invalid result: start %d, end %d", result.Start, result.End)
	case result.End-result.Start+1 != expectedLength:
		return fmt.Errorf("expected length %d, got %d", expectedLength, result.End-result.Start+1)
	}

	previous := a[result.Start]
	for i := result.Start + 1; i <= result.End; i++ {
		if a[i] < previous {
			return fmt.Errorf("element at index %d is smaller than previous element", i)
		}
		previous = a[i]
	}

	return nil
}
