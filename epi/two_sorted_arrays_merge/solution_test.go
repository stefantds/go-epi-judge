package two_sorted_arrays_merge_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/two_sorted_arrays_merge"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]int, int, []int, int)

var solutions = []solutionFunc{
	MergeTwoSortedArrays,
}

func TestMergeTwoSortedArrays(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "two_sorted_arrays_merge.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A              []int
		M              int
		B              []int
		N              int
		ExpectedResult []int
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
			&tc.M,
			&tc.B,
			&tc.N,
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
				result := mergeTwoSortedArraysWrapper(s, tc.A, tc.M, tc.B, tc.N)
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", result, tc.ExpectedResult, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func mergeTwoSortedArraysWrapper(solution solutionFunc, a []int, m int, b []int, n int) []int {
	aCopy := make([]int, len(a))
	_ = copy(aCopy, a)

	bCopy := make([]int, len(b))
	_ = copy(bCopy, b)

	solution(aCopy, m, bCopy, n)
	return aCopy
}
