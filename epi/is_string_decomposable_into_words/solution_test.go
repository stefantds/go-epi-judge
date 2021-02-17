package is_string_decomposable_into_words_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/is_string_decomposable_into_words"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(string, map[string]struct{}) []string

var solutions = []solutionFunc{
	DecomposeIntoDictionaryWords,
}

func TestDecomposeIntoDictionaryWords(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "is_string_decomposable_into_words.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Domain       string
		Dictionary   []string
		Decomposable bool
		Details      string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Domain,
			&tc.Dictionary,
			&tc.Decomposable,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := decomposeIntoDictionaryWordsWrapper(s, tc.Domain, tc.Dictionary, tc.Decomposable); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func decomposeIntoDictionaryWordsWrapper(solution solutionFunc, domain string, dictionary []string, decomposable bool) error {
	dictionaryMap := make(map[string]struct{}, len(dictionary))
	for _, s := range dictionary {
		dictionaryMap[s] = struct{}{}
	}

	result := solution(domain, dictionaryMap)

	if !decomposable {
		if len(result) != 0 {
			return fmt.Errorf("domain is not decomposable. Got: %v", result)
		}
		return nil
	}

	for _, w := range result {
		if _, ok := dictionaryMap[w]; !ok {
			return fmt.Errorf("result uses words not in dictionary. Got: %v", result)
		}
	}

	if strings.Join(result, "") != domain {
		return fmt.Errorf("result is not composed into domain. Domain: %v, got: %v", domain, result)
	}

	return nil
}
