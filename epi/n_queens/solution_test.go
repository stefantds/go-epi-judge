package n_queens_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/n_queens"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int) [][]int

var solutions = []solutionFunc{
	NQueens,
}

func TestNQueens(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "n_queens.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		N              int
		ExpectedResult [][]int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
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
				result := s(tc.N)
				if !equal(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", result, tc.ExpectedResult, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func equal(result, expected [][]int) bool {
	sort.Slice(expected, func(i, j int) bool {
		return utils.LexIntsCompare(expected[i], expected[j])
	})

	sort.Slice(result, func(i, j int) bool {
		return utils.LexIntsCompare(result[i], result[j])
	})

	return reflect.DeepEqual(result, expected)
}
