package sudoku_solve_test

import (
	"errors"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/stefantds/go-epi-judge/utils"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/sudoku_solve"
)

func TestSolveSudoku(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "sudoku_solve.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		PartialAssignment [][]int
		Details           string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.PartialAssignment,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := solveSudokuWrapper(tc.PartialAssignment); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func solveSudokuWrapper(partialAssignment [][]int) error {
	solved := make([][]int, len(partialAssignment))
	for i, row := range partialAssignment {
		solved[i] = make([]int, len(row))
		copy(solved[i], row)
	}

	SolveSudoku(solved)

	if len(solved) != len(partialAssignment) {
		return fmt.Errorf("initial cell assignment has been changed: got:\n%v", utils.MatrixFmt{solved})
	}

	for i, br := range partialAssignment {
		sr := solved[i]

		if len(br) != len(sr) {
			return fmt.Errorf("initial cell assignment has been changed: got:\n%v", utils.MatrixFmt{solved})
		}

		for j := 0; j < len(br); j++ {
			if br[j] != 0 && br[j] != sr[j] {
				return fmt.Errorf("initial cell assignment has been changed: got:\n%v", utils.MatrixFmt{solved})
			}
		}
	}

	blockSize := int(math.Sqrt(float64(len(solved))))

	for i := 0; i < len(solved); i++ {
		if err := assertUniqueSeq(solved[i]); err != nil {
			return fmt.Errorf("%s: got:\n%v", err, utils.MatrixFmt{solved})
		}
		if err := assertUniqueSeq(gatherColumn(solved, i)); err != nil {
			return fmt.Errorf("%s: got:\n%v", err, utils.MatrixFmt{solved})
		}
		if err := assertUniqueSeq(gatherSquareBlock(solved, blockSize, i)); err != nil {
			return fmt.Errorf("%s: got:\n%v", err, utils.MatrixFmt{solved})
		}
	}

	return nil
}

func assertUniqueSeq(seq []int) error {
	seen := make(map[int]bool)
	for _, x := range seq {
		switch {
		case x == 0:
			return errors.New("cell left uninitialized")
		case x < 0, x > len(seq):
			return errors.New("cell value out of range")
		case seen[x]:
			return errors.New("duplicate value in section")
		}
		seen[x] = true
	}
	return nil
}

func gatherColumn(data [][]int, i int) []int {
	result := make([]int, len(data))
	for j, row := range data {
		result[j] = row[i]
	}
	return result
}

func gatherSquareBlock(data [][]int, blockSize, n int) []int {
	result := make([]int, 0)
	blockX := n % blockSize
	blockY := n / blockSize
	for i := blockX * blockSize; i < (blockX+1)*blockSize; i++ {
		for j := blockY * blockSize; j < (blockY+1)*blockSize; j++ {
			result = append(result, data[i][j])
		}
	}

	return result
}
