package search_frequent_items_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/search_frequent_items"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int, <-chan string) []string

var solutions = []solutionFunc{
	SearchFrequentItems,
}

func TestSearchFrequentItems(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "search_frequent_items.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		K              int
		Stream         []string
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
			&tc.K,
			&tc.Stream,
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
				result := searchFrequentItemsWrapper(s, tc.K, tc.Stream)
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

func searchFrequentItemsWrapper(solution solutionFunc, k int, stream []string) []string {
	streamChan := make(chan string, len(stream))
	for _, v := range stream {
		streamChan <- v
	}
	close(streamChan)

	return solution(k, streamChan)
}

func equal(result []string, expected []string) bool {
	sort.Strings(expected)
	sort.Strings(result)
	return reflect.DeepEqual(result, expected)
}
