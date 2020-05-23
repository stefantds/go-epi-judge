package epi_test

import (
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func checkEliminateDuplicate() error {
	//TODO
	return nil
}

func TestEliminateDuplicate(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "remove_duplicates.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Names   []Name
		Details string
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
			EliminateDuplicate(tc.Names)
			err := checkEliminateDuplicate()
			if err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}
