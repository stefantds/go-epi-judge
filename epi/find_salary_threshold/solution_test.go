package find_salary_threshold_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/find_salary_threshold"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int, []int) float64

var solutions = []solutionFunc{
	FindSalaryCap,
}

func TestFindSalaryCap(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "find_salary_threshold.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		TargetPayroll   int
		CurrentSalaries []int
		ExpectedResult  float64
		Details         string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.TargetPayroll,
			&tc.CurrentSalaries,
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
				result := s(tc.TargetPayroll, tc.CurrentSalaries)
				if !utils.EqualFloat(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}
