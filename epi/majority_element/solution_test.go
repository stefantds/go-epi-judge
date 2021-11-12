package majority_element_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/majority_element"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(<-chan string) string

var solutions = []solutionFunc{
	MajoritySearch,
}

func TestMajoritySearch(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "majority_element.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Stream         []string
		ExpectedResult string
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
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
				result := majoritySearchWrapper(s, tc.Stream)
				if result != tc.ExpectedResult {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", result, tc.ExpectedResult, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func majoritySearchWrapper(solution solutionFunc, stream []string) string {
	streamChan := make(chan string, len(stream))
	for _, v := range stream {
		streamChan <- v
	}
	close(streamChan)

	return solution(streamChan)
}
