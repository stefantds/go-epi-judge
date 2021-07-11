package kth_largest_element_in_long_array_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/kth_largest_element_in_long_array"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(<-chan int, int) int

var solutions = []solutionFunc{
	FindKthLargestUnknownLength,
}

func TestFindKthLargestUnknownLength(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "kth_largest_element_in_long_array.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Stream         []int
		K              int
		ExpectedResult int
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
			&tc.K,
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
				result := findKthLargestUnknownLengthWrapper(s, tc.Stream, tc.K)
				if result != tc.ExpectedResult {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findKthLargestUnknownLengthWrapper(solution solutionFunc, stream []int, k int) int {
	streamChan := make(chan int, len(stream))
	for _, s := range stream {
		streamChan <- s
	}
	close(streamChan)

	return solution(streamChan, k)
}
