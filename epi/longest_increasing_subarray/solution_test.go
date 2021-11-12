package longest_increasing_subarray_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/longest_increasing_subarray"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]int) (start, end int)

var solutions = []solutionFunc{
	FindLongestIncreasingSubarray,
}

func TestFindLongestIncreasingSubarray(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "longest_increasing_subarray.tsv")
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

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := findLongestIncreasingSubarrayWrapper(s, tc.A, tc.ExpectedLength); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findLongestIncreasingSubarrayWrapper(solution solutionFunc, a []int, expectedLength int) error {
	start, end := solution(a)

	switch {
	case start < 0 || start >= len(a):
		return fmt.Errorf("invalid start index %d", start)
	case end < 0 || end >= len(a):
		return fmt.Errorf("invalid end index %d", end)
	case end < start:
		return fmt.Errorf("invalid result: start %d, end %d", start, end)
	case end-start+1 != expectedLength:
		return fmt.Errorf("expected length %d, got %d", expectedLength, end-start+1)
	}

	previous := a[start]
	for i := start + 1; i <= end; i++ {
		if a[i] < previous {
			return fmt.Errorf("element at index %d is smaller than previous element", i)
		}
		previous = a[i]
	}

	return nil
}
