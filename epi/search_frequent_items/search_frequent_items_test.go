package search_frequent_items_test

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/search_frequent_items"
)

func TestSearchFrequentItems(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "search_frequent_items.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		K              int
		Stream         []string
		ExpectedResult []string
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.K,
			&tc.Stream,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := SearchFrequentItems(tc.K, tc.Stream)
			if err := compareResult(result, tc.ExpectedResult); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func compareResult(result []string, expected []string) error {
	sort.Strings(expected)
	sort.Strings(result)
	if !reflect.DeepEqual(result, expected) {
		return fmt.Errorf("expected %v, got %v", expected, result)
	}
	return nil
}
