package task_pairing_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/task_pairing"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]Task) [][2]Task

var solutions = []solutionFunc{
	OptimumTaskAssignment,
}

func TestOptimumTaskAssignment(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "task_pairing.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		TaskDurations  []Task
		ExpectedResult [][2]Task
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.TaskDurations,
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
				result := s(tc.TaskDurations)
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

func equal(result, expected [][2]Task) bool {
	for i := 0; i < len(result); i++ {
		sort.Ints(result[i][:])
	}

	for i := 0; i < len(expected); i++ {
		sort.Ints(expected[i][:])
	}

	sort.Slice(expected, func(i, j int) bool {
		return utils.LexIntsCompare(expected[i][:], expected[j][:])
	})

	sort.Slice(result, func(i, j int) bool {
		return utils.LexIntsCompare(result[i][:], result[j][:])
	})

	return reflect.DeepEqual(result, expected)
}
