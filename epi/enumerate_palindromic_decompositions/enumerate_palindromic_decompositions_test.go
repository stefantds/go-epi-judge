package enumerate_palindromic_decompositions_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/enumerate_palindromic_decompositions"
	"github.com/stefantds/go-epi-judge/utils"
)

func TestPalindromeDecompositions(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "enumerate_palindromic_decompositions.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Text           string
		ExpectedResult [][]string
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Text,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := PalindromeDecompositions(tc.Text)
			if !equal(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func equal(result, expected [][]string) bool {
	sort.Slice(expected, func(i, j int) bool {
		return utils.LexStringsCompare(expected[i], expected[j])
	})

	sort.Slice(result, func(i, j int) bool {
		return utils.LexStringsCompare(result[i], result[j])
	})

	return reflect.DeepEqual(expected, result)
}
