package epi_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestCheckFeasible(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "defective_jugs.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Jugs           [][]int
		L              int
		H              int
		ExpectedResult bool
		Details        string
	}

	parser, err := csv.NewParser(file, &csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Jugs,
			&tc.L,
			&tc.H,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := CheckFeasible(buildJugs(tc.Jugs), tc.L, tc.H)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

func buildJugs(rawJugs [][]int) []Jug {
	result := make([]Jug, len(rawJugs))

	for i, jug := range rawJugs {
		result[i] = Jug{
			Low:  jug[0],
			High: jug[1],
		}
	}

	return result
}
