package fibonacci_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/fibonacci"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int) int

var solutions = []solutionFunc{
	Fibonacci,
}

func TestFibonacci(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "fibonacci.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		N              int
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
