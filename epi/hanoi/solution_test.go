package hanoi_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/hanoi"
	"github.com/stefantds/go-epi-judge/utils"
)

func TestComputeTowerHanoi(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "hanoi.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		NumRings int
		Details  string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.NumRings,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			if err := computeTowerHanoiWrapper(tc.NumRings); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func computeTowerHanoiWrapper(numRings int) error {
	pegs := make([]utils.Stack, NumPegs)

	for i := 0; i < NumPegs; i++ {
		pegs[i] = make(utils.Stack, 0, numRings)
	}

	for i := numRings; i >= 1; i-- {
		pegs[0] = pegs[0].Push(i)
	}

	result := ComputeTowerHanoi(numRings)
	for _, operation := range result {
		from := operation[0]
		to := operation[1]
		if !pegs[to].IsEmpty() && pegs[from].Peek().(int) >= pegs[to].Peek().(int) {
			return fmt.Errorf("illegal move from %d to %d", pegs[from].Peek().(int), pegs[to].Peek().(int))
		}
		var top interface{}
		pegs[from], top = pegs[from].Pop()
		pegs[to] = pegs[to].Push(top)
	}

	fullPeg := make(utils.Stack, 0, numRings)
	for i := numRings; i >= 1; i-- {
		fullPeg = fullPeg.Push(i)
	}

	expectedPegs1 := []utils.Stack{
		{},
		fullPeg,
		{},
	}

	expectedPegs2 := []utils.Stack{
		{},
		{},
		fullPeg,
	}

	// check if the result is one of the accepted pegs configs
	if !reflect.DeepEqual(pegs, expectedPegs1) &&
		!reflect.DeepEqual(pegs, expectedPegs2) {
		return fmt.Errorf("pegs are not in the expected configuration")
	}

	return nil
}
