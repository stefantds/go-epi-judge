package epi_test

import (
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestFindAmpleCity(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "refueling_schedule.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Gallons   []int
		Distances []int
		Details   string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Gallons,
			&tc.Distances,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := findAmpleCityWrapper(tc.Gallons, tc.Distances); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findAmpleCityWrapper(gallons []int, distances []int) error {
	result := FindAmpleCity(gallons, distances)
	numCities := len(gallons)
	tank := 0

	for i := 0; i < numCities; i++ {
		city := (result + i) % numCities
		tank += gallons[city]*MPG - distances[city]

		if tank < 0 {
			return fmt.Errorf("out of gas on city %d", city)
		}
	}

	return nil
}
