package epi_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestStackWithMax(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "stack_with_max.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Operations stackWithMaxDecoder
		Details    string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := stackWithMaxTester(tc.Operations.Value); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func stackWithMaxTester(operations []*StackWithMaxOp) error {
	var q StackWithMax
	for opIdx, o := range operations {
		switch o.Op {
		case "Stack":
			q = NewStackWithMax()
		case "push":
			q.Push(o.Arg)
		case "pop":
			result := q.Pop()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.Arg, result)
			}
		case "max":
			result := q.Max()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.Arg, result)
			}
		case "empty":
			var result int
			if q.Empty() {
				result = 1
			} else {
				result = 0
			}
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.Arg, result)
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

func (o *stackWithMaxDecoder) DecodeRecord(record string) error {
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
