package search_for_missing_element_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/search_for_missing_element"
)

func TestFindDuplicateMissing(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "find_missing_and_duplicate.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A              []int
		ExpectedResult []int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.A,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			resultDupl, resultMissing := FindDuplicateMissing(tc.A)
			if !reflect.DeepEqual([]int{resultDupl, resultMissing}, tc.ExpectedResult) {
				t.Errorf("\nexpected:\n%v\ngot:\n%v", tc.ExpectedResult, []int{resultDupl, resultMissing})
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}
