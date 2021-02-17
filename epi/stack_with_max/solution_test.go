package stack_with_max_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/stack_with_max"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func() Solution

var solutions = []solutionFunc{
	NewStackWithMax,
}

func TestStackWithMax(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "stack_with_max.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Operations stackWithMaxDecoder
		Details    string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Operations,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := stackWithMaxTester(s, tc.Operations.Value); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func stackWithMaxTester(sol solutionFunc, operations []*StackWithMaxOp) error {
	var q Solution
	for opIdx, o := range operations {
		switch o.Op {
		case "Stack":
			q = sol()
		case "push":
			q.Push(o.Arg)
		case "pop":
			result := q.Pop()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: got: %v, want: %v", opIdx, o.Op, result, o.Arg)
			}
		case "max":
			result := q.Max()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: got: %v, want: %v", opIdx, o.Op, result, o.Arg)
			}
		case "empty":
			var result int
			if q.Empty() {
				result = 1
			} else {
				result = 0
			}
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: got: %v, want: %v", opIdx, o.Op, result, o.Arg)
			}
		}
	}

	return nil
}

type StackWithMaxOp struct {
	Op  string
	Arg int
}

type stackWithMaxDecoder struct {
	Value []*StackWithMaxOp
}

func (o *stackWithMaxDecoder) DecodeField(record string) error {
	allData := make([][3]interface{}, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([]*StackWithMaxOp, len(allData))
	for i := 0; i < len(allData); i++ {
		result[i] = &StackWithMaxOp{
			Op:  allData[i][0].(string),
			Arg: int(allData[i][1].(float64)),
		}
	}

	o.Value = result
	return nil
}
