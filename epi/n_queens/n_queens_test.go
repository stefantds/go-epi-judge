package n_queens_test

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/n_queens"
	"github.com/stefantds/go-epi-judge/utils"
)

func TestNQueens(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "n_queens.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		N              int
		ExpectedResult [][]int
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.N,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := NQueens(tc.N)
			if !equal(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func equal(result, expected [][]int) bool {
	sort.Slice(expected, func(i, j int) bool {
		return utils.LexIntsCompare(expected[i], expected[j])
	})

	sort.Slice(result, func(i, j int) bool {
		return utils.LexIntsCompare(result[i], result[j])
	})

	return reflect.DeepEqual(result, expected)
}
