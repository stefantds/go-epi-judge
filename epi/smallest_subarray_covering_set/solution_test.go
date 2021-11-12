package smallest_subarray_covering_set_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/smallest_subarray_covering_set"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]string, map[string]struct{}) (int, int)

var solutions = []solutionFunc{
	FindSmallestSubarrayCoveringSet,
}

func TestFindSmallestSubarrayCoveringSet(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "smallest_subarray_covering_set.tsv")
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

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := findSmallestSubarrayCoveringSetWrapper(s, tc.Paragraph, tc.Keywords, tc.ExpectedLength); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findSmallestSubarrayCoveringSetWrapper(solution solutionFunc, paragraph []string, keywords []string, expectedLength int) error {
	set := make(map[string]struct{}, len(keywords))
	copySet := make(map[string]struct{}, len(keywords))
	for _, s := range keywords {
		set[s] = struct{}{}
		copySet[s] = struct{}{}
	}

	start, end := solution(paragraph, set)

	switch {
	case start < 0 || start >= len(paragraph):
		return fmt.Errorf("invalid start index %d", start)
	case end < 0 || end >= len(paragraph):
		return fmt.Errorf("invalid end index %d", end)
	case end < start:
		return fmt.Errorf("invalid result: start %d, end %d", start, end)
	case end-start+1 != expectedLength:
		return fmt.Errorf("expected length %d, got %d", expectedLength, end-start+1)
	}

	for i := start; i <= end; i++ {
		delete(copySet, paragraph[i])
	}

	if len(copySet) > 0 {
		return errors.New("not all keywords are in the range")
	}

	return nil
}
