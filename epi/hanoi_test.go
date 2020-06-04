package epi_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
	"github.com/stefantds/go-epi-judge/stack"
)

func TestComputeTowerHanoi(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "hanoi.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		NumRings int
		Details  string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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
	pegs := make([]stack.Stack, NumPegs)

	for i := 0; i < NumPegs; i++ {
		pegs[i] = make(stack.Stack, 0, numRings)
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

	fullPeg := make(stack.Stack, 0, numRings)
	for i := numRings; i >= 1; i-- {
		fullPeg = fullPeg.Push(i)
	}

	expectedPegs1 := []stack.Stack{
		{},
		fullPeg,
		{},
	}

	expectedPegs2 := []stack.Stack{
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
