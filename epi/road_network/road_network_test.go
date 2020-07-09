package road_network_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/road_network"
)

func TestFindBestProposals(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "road_network.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		H              highwaySectionsDecoder
		P              highwaySectionsDecoder
		N              int
		ExpectedResult highwaySectionDecoder
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.H,
			&tc.P,
			&tc.N,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := FindBestProposals(tc.H.Values, tc.P.Values, tc.N)
			if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult.Value, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type highwaySectionsDecoder struct {
	Values []HighwaySection
}

func (o *highwaySectionsDecoder) DecodeRecord(record string) error {
	allData := make([][3]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]HighwaySection, len(allData))
	for i := 0; i < len(allData); i++ {
		values[i].X = allData[i][0]
		values[i].Y = allData[i][1]
		values[i].Distance = allData[i][2]
	}

	o.Values = values
	return nil
}

type highwaySectionDecoder struct {
	Value HighwaySection
}

func (o *highwaySectionDecoder) DecodeRecord(record string) error {
	allData := make([]int, 3)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	o.Value.X = allData[0]
	o.Value.Y = allData[1]
	o.Value.Distance = allData[2]

	return nil
}
