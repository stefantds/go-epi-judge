package dutch_national_flag_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/dutch_national_flag"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int, []Color)

var solutions = []solutionFunc{
	DutchFlagPartition,
}

func TestDutchFlagPartition(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "dutch_national_flag.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		PivotIndex int
		A          []Color
		Details    string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.A,
			&tc.PivotIndex,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := dutchFlagPartitionWrapper(s, tc.A, tc.PivotIndex); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func dutchFlagPartitionWrapper(solution solutionFunc, a []Color, pivotIdx int) error {
	result := make([]Color, len(a))
	copy(result, a)

	solution(pivotIdx, result)

	count := make(map[Color]int, 3)

	for _, c := range a {
		count[c] += 1
	}

	pivot := a[pivotIdx]

	i := 0
	for i < len(result) && int(result[i]) < int(pivot) {
		count[result[i]]--
		i++
	}

	for i < len(result) && int(result[i]) == int(pivot) {
		count[result[i]]--
		i++
	}

	for i < len(result) && int(result[i]) > int(pivot) {
		count[result[i]]--
		i++
	}

	if i != len(result) {
		return fmt.Errorf("not partitioned after %v-th element", i)
	}
	if count[Red] != 0 || count[White] != 0 || count[Blue] != 0 {
		return errors.New("some elements are missing from original array")
	}

	return nil
}
