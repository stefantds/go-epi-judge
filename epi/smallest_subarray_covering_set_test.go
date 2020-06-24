package epi_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestFindSmallestSubarrayCoveringSet(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "smallest_subarray_covering_set.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Paragraph      []string
		Keywords       []string
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
			&tc.Paragraph,
			&tc.Keywords,
			&tc.ExpectedLength,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := findSmallestSubarrayCoveringSetWrapper(tc.Paragraph, tc.Keywords, tc.ExpectedLength); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findSmallestSubarrayCoveringSetWrapper(paragraph []string, keywords []string, expectedLength int) error {
	set := make(map[string]struct{}, len(keywords))
	copySet := make(map[string]struct{}, len(keywords))
	for _, s := range keywords {
		set[s] = struct{}{}
		copySet[s] = struct{}{}
	}

	result := FindSmallestSubarrayCoveringSet(paragraph, set)

	switch {
	case result.Start < 0 || result.Start >= len(paragraph):
		return fmt.Errorf("invalid start index %d", result.Start)
	case result.End < 0 || result.End >= len(paragraph):
		return fmt.Errorf("invalid end index %d", result.End)
	case result.End < result.Start:
		return fmt.Errorf("invalid result: start %d, end %d", result.Start, result.End)
	case result.End-result.Start+1 != expectedLength:
		return fmt.Errorf("expected length %d, got %d", expectedLength, result.End-result.Start+1)
	}

	for i := result.Start; i <= result.End; i++ {
		delete(copySet, paragraph[i])
	}

	if len(copySet) > 0 {
		return errors.New("not all keywords are in the range")
	}

	return nil
}
