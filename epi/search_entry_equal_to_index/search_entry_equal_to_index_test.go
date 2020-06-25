package search_entry_equal_to_index_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/search_entry_equal_to_index"
)

func TestSearchEntryEqualToItsIndex(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "search_entry_equal_to_index.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A       []int
		Details string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.A,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := searchEntryEqualToItsIndexWrapper(tc.A); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func searchEntryEqualToItsIndexWrapper(a []int) error {
	result := SearchEntryEqualToItsIndex(a)

	if result < -1 || result > len(a)-1 {
		return fmt.Errorf("invalid index %d", result)
	}
	if result != -1 {
		if a[result] != result {
			return fmt.Errorf("got index %d; a[%d] is %d", result, result, a[result])
		}
	} else {
		for i, x := range a {
			if i == x {
				return errors.New("got -1 but there are entries which equal to their index")
			}
		}
	}

	return nil
}
