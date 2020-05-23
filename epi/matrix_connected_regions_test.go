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

func TestFlipColor(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "painting.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		X              int
		Y              int
		Image          imageDecoder
		ExpectedResult imageDecoder
		Details        string
	}

	parser, err := csv.NewParser(file, &csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.X,
			&tc.Y,
			&tc.Image,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			FlipColor(tc.X, tc.Y, tc.Image.Value)
			if !reflect.DeepEqual(tc.Image.Value, tc.ExpectedResult.Value) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult.Value, tc.Image.Value)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

type imageDecoder struct {
	Value [][]bool
}

func (d *imageDecoder) DecodeRecord(record string) error {
	allData := make([][]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([][]bool, len(allData))
	for i := 0; i < len(allData); i++ {
		result[i] = make([]bool, len(allData[0]))
		for j := 0; j < len(allData[0]); j++ {
			switch allData[i][j] {
			case 0:
				result[i][j] = false
			case 1:
				result[i][j] = true
			}
		}
	}

	d.Value = result
	return nil
}
