package smallest_subarray_covering_all_values_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/smallest_subarray_covering_all_values"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]string, []string) (int, int)

var solutions = []solutionFunc{
	FindSmallestSequentiallyCoveringSubset,
}

func TestFindSmallestSequentiallyCoveringSubset(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "smallest_subarray_covering_all_values.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Paragraph      []string
		Keywords       []string
		ExpectedResult int
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
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result, err := findSmallestSequentiallyCoveringSubsettWrapper(s, tc.Paragraph, tc.Keywords)
				if err != nil {
					t.Fatal(err)
				}
				if result != tc.ExpectedResult {
					t.Errorf("expected min length %v, got %v", tc.ExpectedResult, result)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findSmallestSequentiallyCoveringSubsettWrapper(solution solutionFunc, paragraph, keywords []string) (int, error) {
	start, end := solution(paragraph, keywords)
	if start < 0 {
		return 0, errors.New("subarray start index is negative")
	}

	kwIdx := 0
	paraIdx := start

	for kwIdx < len(keywords) {
		if paraIdx > end {
			return 0, errors.New("not all keywords are in the generated subarray")
		}
		if paraIdx >= len(paragraph) {
			return 0, errors.New("subarray end index exceeds array size")
		}
		if paragraph[paraIdx] == keywords[kwIdx] {
			kwIdx++
		}
		paraIdx++
	}
	return end - start + 1, nil
}
