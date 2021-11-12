package matrix_rotation_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/matrix_rotation"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([][]int)

var solutions = []solutionFunc{
	RotateMatrix,
}

func TestRotateMatrix(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "matrix_rotation.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		SquareMatrix   [][]int
		ExpectedResult [][]int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.SquareMatrix,
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
				result := rotateMatrixWrapper(s, tc.SquareMatrix)
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", utils.MatrixFormatter(result), utils.MatrixFormatter(tc.ExpectedResult), tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func rotateMatrixWrapper(solution solutionFunc, squareMatrix [][]int) [][]int {
	result := make([][]int, len(squareMatrix))
	for i := range squareMatrix {
		result[i] = make([]int, len(squareMatrix[i]))
		copy(result[i], squareMatrix[i])
	}

	solution(result)
	return result
}
