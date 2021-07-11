package sunset_view_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/sunset_view"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(<-chan int) []int

var solutions = []solutionFunc{
	ExamineBuildingsWithSunset,
}

func TestExamineBuildingsWithSunset(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "sunset_view.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Sequence       []int
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
				result := examineBuildingsWithSunsetWrapper(s, tc.Sequence)
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func examineBuildingsWithSunsetWrapper(solution solutionFunc, sequence []int) []int {
	sequenceChan := make(chan int, len(sequence))
	for _, v := range sequence {
		sequenceChan <- v
	}
	close(sequenceChan)

	return solution(sequenceChan)
}
