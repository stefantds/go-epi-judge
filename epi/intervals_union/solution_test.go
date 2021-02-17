package intervals_union_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/intervals_union"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]Interval) []Interval

var solutions = []solutionFunc{
	UnionOfIntervals,
}

func TestUnionOfIntervals(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "intervals_union.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Intervals      intervalsDecoder
		ExpectedResult intervalsDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Intervals,
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
				result := s(tc.Intervals.Value)
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

type intervalsDecoder struct {
	Value []Interval
}

func (o *intervalsDecoder) DecodeField(record string) error {
	allData := make([][4]interface{}, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]Interval, len(allData))
	for i := 0; i < len(allData); i++ {
		leftVal := int(allData[i][0].(float64))
		leftIsClosed := allData[i][1].(bool)
		rightVal := int(allData[i][2].(float64))
		rightIsClosed := allData[i][3].(bool)

		values[i] = Interval{
			Left: Endpoint{
				Val:      leftVal,
				IsClosed: leftIsClosed,
			},
			Right: Endpoint{rightVal, rightIsClosed},
		}
	}

	o.Value = values
	return nil
}
