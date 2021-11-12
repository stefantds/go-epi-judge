package road_network_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/road_network"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]HighwaySection, []HighwaySection, int) *HighwaySection

var solutions = []solutionFunc{
	FindBestProposals,
}

func TestFindBestProposals(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "road_network.tsv")
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

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result := s(tc.H.Values, tc.P.Values, tc.N)
				if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", result, tc.ExpectedResult.Value, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type highwaySectionsDecoder struct {
	Values []HighwaySection
}

func (o *highwaySectionsDecoder) DecodeField(record string) error {
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

func (o *highwaySectionDecoder) DecodeField(record string) error {
	allData := make([]int, 3)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	o.Value.X = allData[0]
	o.Value.Y = allData[1]
	o.Value.Distance = allData[2]

	return nil
}
