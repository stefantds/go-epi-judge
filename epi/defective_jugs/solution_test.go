package defective_jugs_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/defective_jugs"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]Jug, int, int) bool

var solutions = []solutionFunc{
	CheckFeasible,
}

func TestCheckFeasible(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "defective_jugs.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Jugs           [][]int
		L              int
		H              int
		ExpectedResult bool
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Jugs,
			&tc.L,
			&tc.H,
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
				result := CheckFeasible(buildJugs(tc.Jugs), tc.L, tc.H)
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

func buildJugs(rawJugs [][]int) []Jug {
	result := make([]Jug, len(rawJugs))

	for i, jug := range rawJugs {
		result[i] = Jug{
			Low:  jug[0],
			High: jug[1],
		}
	}

	return result
}
