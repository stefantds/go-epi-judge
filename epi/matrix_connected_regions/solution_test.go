package matrix_connected_regions_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/matrix_connected_regions"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int, int, [][]bool)

var solutions = []solutionFunc{
	FlipColor,
}

func TestFlipColor(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "painting.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		X              int
		Y              int
		Image          imageDecoder
		ExpectedResult imageDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.X,
			&tc.Y,
			&tc.Image,
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
				result := flipColorWrapper(s, tc.X, tc.Y, tc.Image.Value)
				if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", utils.MatrixFormatter(result), utils.MatrixFormatter(tc.ExpectedResult.Value))
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type imageDecoder struct {
	Value [][]bool
}

func (d *imageDecoder) DecodeField(record string) error {
	allData := make([][]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([][]bool, len(allData))
	for i := 0; i < len(allData); i++ {
		result[i] = make([]bool, len(allData[0]))
		for j := 0; j < len(allData[0]); j++ {
			switch allData[i][j] {
			case 0:
				result[i][j] = false
			case 1:
				result[i][j] = true
			}
		}
	}

	d.Value = result
	return nil
}

func flipColorWrapper(solution solutionFunc, x int, y int, image [][]bool) [][]bool {
	solution(x, y, image)
	return image
}
