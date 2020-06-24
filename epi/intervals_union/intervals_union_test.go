package intervals_union_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/intervals_union"
)

func TestUnionOfIntervals(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "intervals_union.tsv"
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

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := UnionOfIntervals(tc.Intervals.Value)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult.Value, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type intervalsDecoder struct {
	Value []Interval
}

func (o *intervalsDecoder) DecodeRecord(record string) error {
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
