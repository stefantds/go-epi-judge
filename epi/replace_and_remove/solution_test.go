package replace_and_remove_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/replace_and_remove"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int, []rune) int

var solutions = []solutionFunc{
	ReplaceAndRemove,
}

func TestReplaceAndRemove(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "replace_and_remove.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Size           int
		S              []string
		ExpectedResult []string
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Size,
			&tc.S,
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
				result, err := replaceAndRemoveWrapper(s, tc.Size, tc.S)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", result, tc.ExpectedResult, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func replaceAndRemoveWrapper(solution solutionFunc, size int, s []string) ([]string, error) {
	allChars := make([]rune, len(s))
	for i, c := range []rune(strings.Join(s, "")) {
		allChars[i] = c
	}
	resSize := solution(size, allChars)

	if resSize > len(s) {
		return nil, errors.New("result can't be greater than the original size")
	}
	result := make([]string, resSize)
	for i := 0; i < resSize; i++ {
		result[i] = string(allChars[i])
	}

	return result, nil
}
