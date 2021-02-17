package anagrams_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/anagrams"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(dictionary []string) [][]string

var solutions = []solutionFunc{
	FindAnagrams,
}

func TestFindAnagrams(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "anagrams.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Dictionary     []string
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
			&tc.Dictionary,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result := s(tc.Dictionary)
				if !equal(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func equal(result, expected [][]string) bool {
	for _, l := range expected {
		sort.Strings(l)
	}

	sort.Slice(expected, func(i, j int) bool {
		return utils.LexStringsCompare(expected[i], expected[j])
	})

	for _, l := range result {
		sort.Strings(l)
	}

	sort.Slice(result, func(i, j int) bool {
		return utils.LexStringsCompare(result[i], result[j])
	})

	return reflect.DeepEqual(expected, result)
}
