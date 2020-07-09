package max_of_sliding_window_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/max_of_sliding_window"
)

func TestComputeTrafficVolumes(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "max_of_sliding_window.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A              trafficElementsDecoder
		W              int
		ExpectedResult trafficElementsDecoder
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.A,
			&tc.W,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := ComputeTrafficVolumes(tc.A.Values, tc.W)
			if !reflect.DeepEqual(result, tc.ExpectedResult.Values) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult.Values, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type trafficElementsDecoder struct {
	Values []TrafficElement
}

func (o *trafficElementsDecoder) DecodeRecord(record string) error {
	allData := make([][2]float64, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]TrafficElement, len(allData))
	for i := 0; i < len(allData); i++ {
		values[i].Time = int(allData[i][0])
		values[i].Volume = allData[i][1]
	}

	o.Values = values
	return nil
}
