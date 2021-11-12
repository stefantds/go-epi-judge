package max_of_sliding_window_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/max_of_sliding_window"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]TrafficElement, int) []TrafficElement

var solutions = []solutionFunc{
	ComputeTrafficVolumes,
}

func TestComputeTrafficVolumes(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "max_of_sliding_window.tsv")
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

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result := s(tc.A.Values, tc.W)
				if !reflect.DeepEqual(result, tc.ExpectedResult.Values) {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", result, tc.ExpectedResult.Values, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type trafficElementsDecoder struct {
	Values []TrafficElement
}

func (o *trafficElementsDecoder) DecodeField(record string) error {
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
