package epi_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestFillSurroundedRegions(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "matrix_enclosed_regions.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Board          boardDecoder
		ExpectedResult boardDecoder
		Details        string
	}

	parser, err := csv.NewParser(file, &csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Board,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			FillSurroundedRegions(tc.Board.Value)
			if !reflect.DeepEqual(tc.Board.Value, tc.ExpectedResult.Value) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, tc.Board.Value)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

type boardDecoder struct {
	Value [][]rune
}

func (d *boardDecoder) DecodeRecord(record string) error {
	allData := make([][]string, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([][]rune, len(allData))
	for i := 0; i < len(allData); i++ {
		result[i] = make([]rune, len(allData[0]))
		for j := 0; j < len(allData[0]); j++ {
			result[i][j] = []rune(allData[i][j])[0]
		}
	}

	d.Value = result
	return nil
}
