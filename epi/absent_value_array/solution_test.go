package absent_value_array_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/absent_value_array"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(stream chan int) int

var solutions = []solutionFunc{
	FindMissingElement,
}

func TestFindMissingElement(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "absent_value_array.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Stream  []int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Stream,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				err := findMissingElementWrapper(s, tc.Stream)
				if err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findMissingElementWrapper(solution solutionFunc, stream []int) error {
	streamChan := make(chan int, len(stream))
	for _, v := range stream {
		streamChan <- v
	}
	close(streamChan)

	res := solution(streamChan)

	for _, i := range stream {
		if i == res {
			return fmt.Errorf("%d appears in stream", res)
		}
	}

	return nil
}
