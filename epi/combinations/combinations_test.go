package combinations_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/go-epi-judge/utils"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/combinations"
)

func TestCombinations(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "combinations.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		N              int
		K              int
		ExpectedResult [][]int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.N,
			&tc.K,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := Combinations(tc.N, tc.K)
			if !equal(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
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
