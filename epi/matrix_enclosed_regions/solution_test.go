package matrix_enclosed_regions_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/matrix_enclosed_regions"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([][]Color)

var solutions = []solutionFunc{
	FillSurroundedRegions,
}

func TestFillSurroundedRegions(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "matrix_enclosed_regions.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Board          [][]Color
		ExpectedResult [][]Color
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Board,
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
				result := fillSurroundedRegionsWrapper(s, tc.Board)
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", utils.MatrixFormatter(result), utils.MatrixFormatter(tc.ExpectedResult))
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func fillSurroundedRegionsWrapper(solution solutionFunc, board [][]Color) [][]Color {
	result := make([][]Color, len(board))
	for i := range board {
		result[i] = make([]Color, len(board[i]))
		copy(result[i], board[i])
	}

	solution(result)
	return result
}
