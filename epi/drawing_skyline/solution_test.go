package drawing_skyline_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/drawing_skyline"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]Rect) []Rect

var solutions = []solutionFunc{
	DrawingSkylines,
}

func TestDrawingSkylines(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "drawing_skyline.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Buildings      rectsDecoder
		ExpectedResult rectsDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Buildings,
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
				result := s(tc.Buildings.Values)
				if !reflect.DeepEqual(result, tc.ExpectedResult.Values) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult.Values)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type rectsDecoder struct {
	Values []Rect
}

func (o *rectsDecoder) DecodeField(record string) error {
	allData := make([][3]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]Rect, len(allData))
	for i := 0; i < len(allData); i++ {
		values[i].Left = allData[i][0]
		values[i].Right = allData[i][1]
		values[i].Height = allData[i][2]
	}

	o.Values = values
	return nil
}
