package epi_test

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestDecomposeIntoDictionaryWords(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "is_string_decomposable_into_words.tsv"
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

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := decomposeIntoDictionaryWordsWrapper(tc.Domain, tc.Dictionary, tc.Decomposable); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func decomposeIntoDictionaryWordsWrapper(domain string, dictionary []string, decomposable bool) error {
	dictionaryMap := make(map[string]struct{}, len(dictionary))
	for _, s := range dictionary {
		dictionaryMap[s] = struct{}{}
	}

	result := DecomposeIntoDictionaryWords(domain, dictionaryMap)

	if !decomposable {
		if len(result) != 0 {
			return errors.New("domain is not decomposable")
		}
		return nil
	}

	for _, w := range result {
		if _, ok := dictionaryMap[w]; !ok {
			return errors.New("result uses words not in dictionary")
		}
	}

	if strings.Join(result, "") != domain {
		return errors.New("result is not composed into domain")
	}

	return nil
}
