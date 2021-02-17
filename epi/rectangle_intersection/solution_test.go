package rectangle_intersection_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/rectangle_intersection"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(Rect, Rect) Rect

var solutions = []solutionFunc{
	IntersectRectangle,
}

func TestIntersectRectangle(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "rectangle_intersection.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		R1             rectDecoder
		R2             rectDecoder
		ExpectedResult rectDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.R1,
			&tc.R2,
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
				result := s(tc.R1.Value, tc.R2.Value)
				if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult.Value)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type rectDecoder struct {
	Value Rect
}

func (i *rectDecoder) DecodeField(record string) error {
	var allData [4]int
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	i.Value.X = allData[0]
	i.Value.Y = allData[1]
	i.Value.Width = allData[2]
	i.Value.Height = allData[3]

	return nil
}
