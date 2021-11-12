package sort_almost_sorted_array_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/sort_almost_sorted_array"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(<-chan int, int) []int

var solutions = []solutionFunc{
	SortApproximatelySortedData,
}

func TestSortApproximatelySortedData(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "sort_almost_sorted_array.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Sequence       []int
		K              int
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
			&tc.Sequence,
			&tc.K,
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
				result := sortApproximatelySortedDataWrapper(s, tc.Sequence, tc.K)
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

func sortApproximatelySortedDataWrapper(solution solutionFunc, sequence []int, k int) []int {
	sequenceChan := make(chan int, len(sequence))
	for _, v := range sequence {
		sequenceChan <- v
	}
	close(sequenceChan)

	return solution(sequenceChan, k)
}
