package epi_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/goepijudge/epi"
)

func TestComputeTrafficVolumes(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "max_of_sliding_window.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A              []TrafficElement
		W              int
		ExpectedResult []TrafficElement
		Details        string
	}

	parser, err := csv.NewParser(file, &csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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
			result := ComputeTrafficVolumes(tc.A, tc.W)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}
