package k_closest_stars_test

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/k_closest_stars"
)

func TestFindClosestKStars(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "k_closest_stars.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Stars          starsDecoder
		K              int
		ExpectedResult []float64
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Stars,
			&tc.K,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := FindClosestKStars(tc.Stars.Value, tc.K)
			if !equal(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func equal(result []Star, expected []float64) bool {
	if len(expected) != len(result) {
		return false
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Distance() < result[j].Distance() })

	for i := 0; i < len(result); i++ {
		if result[i].Distance() != expected[i] {
			return false
		}
	}

	return true
}

type starsDecoder struct {
	Value []Star
}

func (d *starsDecoder) DecodeField(record string) error {
	allData := make([][3]float64, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	d.Value = make([]Star, len(allData))
	for i := 0; i < len(allData); i++ {
		d.Value[i] = Star{
			X: allData[i][0],
			Y: allData[i][1],
			Z: allData[i][2],
		}
	}

	return nil
}
