package minimum_points_covering_intervals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/minimum_points_covering_intervals"
)

func TestFindMinimumVisits(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "minimum_points_covering_intervals.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Intervals      intervalsDecoder
		ExpectedResult int
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := FindMinimumVisits(tc.Intervals.Values)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type intervalsDecoder struct {
	Values []Interval
}

func (o *intervalsDecoder) DecodeField(record string) error {
	allData := make([][2]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]Interval, len(allData))
	for i := 0; i < len(allData); i++ {
		values[i].Left = allData[i][0]
		values[i].Right = allData[i][1]
	}

	o.Values = values
	return nil
}
