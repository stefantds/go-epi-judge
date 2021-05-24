package gray_code_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/gray_code"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int) []int64

var solutions = []solutionFunc{
	GrayCode,
}

func TestGrayCode(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "gray_code.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		NumBits int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.NumBits,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := grayCodeWrapper(s, tc.NumBits); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func grayCodeWrapper(solution solutionFunc, numBits int) error {
	result := solution(numBits)
	uniqueEntries := make(map[int64]bool)

	expectedSize := 1 << numBits
	if expectedSize != len(result) {
		return fmt.Errorf("length mismatch: got %d, want %d", len(result), expectedSize)
	}

	for i := 1; i < len(result); i++ {
		if !differsByOneBit(result[i-1], result[i]) {
			if result[i-1] == result[i] {
				return fmt.Errorf("two consecutive entries are equal at index %d and %d", i-1, i)
			}
			return fmt.Errorf("two consecutive entries differ by more than 1 at index %d and %d", i-1, i)
		}
		if _, ok := uniqueEntries[result[i]]; ok == true {
			return fmt.Errorf("not all entries are distint: %d is duplicated", result[i])
		}
		uniqueEntries[result[i]] = true
	}

	return nil
}

func differsByOneBit(x, y int64) bool {
	bitDiff := x ^ y
	return bitDiff != 0 && bitDiff&(bitDiff-1) == 0
}
