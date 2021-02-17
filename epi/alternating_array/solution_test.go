package alternating_array_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/alternating_array"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(a []int)

var solutions = []solutionFunc{
	Rearrange,
}

func TestRearrange(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "alternating_array.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A       []int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.A,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				Rearrange(tc.A)
				if err := rearrangeWrapper(s, tc.A); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func rearrangeWrapper(solution solutionFunc, a []int) error {
	result := make([]int, len(a))
	_ = copy(result, a)

	solution(result)

	if err := utils.AssertAllValuesPresent(a, result); err != nil {
		return err
	}

	return checkOrder(result)
}

func checkOrder(a []int) error {
	for i := 0; i < len(a); i++ {
		if (i % 2) != 0 {
			if a[i] < a[i-1] {
				return fmt.Errorf("wrong order found: got: %d > %d; want: a[%d] <= a[%d]", a[i-1], a[i], i-1, i)
			}
			if i < len(a)-1 {
				if a[i] < a[i+1] {
					return fmt.Errorf("wrong order found: got: %d < %d; want: a[%d] >= a[%d]", a[i], a[i+1], i, i+1)
				}
			}
		} else {
			if i > 0 {
				if a[i-1] < a[i] {
					return fmt.Errorf("wrong order found: got: %d < %d; want: a[%d] >= a[%d]", a[i-1], a[i], i-1, i)
				}
			}
			if i < len(a)-1 {
				if a[i+1] < a[i] {
					return fmt.Errorf("wrong order found: got: %d > %d; want: a[%d] <= a[%d]", a[i], a[i+1], i, i+1)
				}
			}
		}
	}

	return nil
}
