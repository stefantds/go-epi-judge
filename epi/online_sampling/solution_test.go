package online_sampling_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/online_sampling"
	utils "github.com/stefantds/go-epi-judge/test_utils"
	"github.com/stefantds/go-epi-judge/test_utils/stats"
)

type solutionFunc = func(<-chan int, int) []int

var solutions = []solutionFunc{
	OnlineRandomSample,
}

func TestOnlineRandomSample(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "online_sampling.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Stream  []int
		K       int
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
			&tc.K,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := onlineRandomSampleWrapper(s, tc.Stream, tc.K); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func onlineRandomSampleWrapper(solution solutionFunc, stream []int, k int) error {
	return stats.RunFuncWithRetries(
		func() bool {
			return onlineRandomSampleRunner(solution, stream, k)
		},
		errors.New("the results don't match the expected distribution"),
	)
}

func onlineRandomSampleRunner(solution solutionFunc, stream []int, k int) bool {
	const nbRuns = 1000000

	results := make([][]int, nbRuns)
	for i := 0; i < nbRuns; i++ {
		streamChan := make(chan int, len(stream))
		for _, v := range stream {
			streamChan <- v
		}
		close(streamChan)
		results[i] = solution(streamChan, k)
	}

	totalPossibleOutcomes := stats.BinomialCoefficient(len(stream), k)

	combinations := make([][]int, totalPossibleOutcomes)
	for i := 0; i < totalPossibleOutcomes; i++ {
		combinations[i] = stats.ComputeCombinationIdx(stream, k, i)
	}

	sort.Slice(combinations, func(i, j int) bool {
		return utils.LexIntsCompare(combinations[i], combinations[j])
	})

	sequence := make([]int, nbRuns)
	for i, r := range results {
		sort.Ints(r)
		pos := sort.Search(
			len(combinations),
			func(i int) bool { return !utils.LexIntsCompare(combinations[i], r) },
		)
		if pos < len(combinations) && reflect.DeepEqual(combinations[pos], r) {
			sequence[i] = pos
		} else {
			panic("result not in known combinations")
		}
	}

	return stats.CheckSequenceIsUniformlyRandom(sequence, totalPossibleOutcomes, 0.01)
}
