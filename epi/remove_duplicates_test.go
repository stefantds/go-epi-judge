package epi_test

import (
	"fmt"
	"os"
	"sort"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestEliminateDuplicate(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "remove_duplicates.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Names          [][2]string
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
			&tc.Names,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := namesComparator(tc.ExpectedResult,
				eliminateDuplicateWrapper(decodeNames(tc.Names)),
			); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func eliminateDuplicateWrapper(names []Name) []Name {
	EliminateDuplicate(names)
	return names
}

func decodeNames(names [][2]string) []Name {
	result := make([]Name, len(names))

	for i, n := range names {
		result[i].FirstName = n[0]
		result[i].LastName = n[1]
	}

	return result
}

func namesComparator(expected []string, result []Name) error {
	if len(expected) != len(result) {
		return fmt.Errorf("expected %d names in result, got %d", len(expected), len(result))
	}

	sort.SliceStable(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].FirstName < result[j].FirstName
	})

	for i := 0; i < len(result); i++ {
		if expected[i] != result[i].FirstName {
			return fmt.Errorf("unexpected first name found: %s", result[i].FirstName)
		}
	}

	return nil
}
