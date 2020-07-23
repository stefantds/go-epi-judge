package queue_with_max_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/queue_with_max"
)

func TestQueueWithMax(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "queue_with_max.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Operations queueWithMaxDecoder
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := queueWithMaxTester(tc.Operations.Value); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func queueWithMaxTester(operations []*QueueWithMaxOp) error {
	var q QueueWithMax
	for opIdx, o := range operations {
		switch o.Op {
		case "QueueWithMax":
			q = NewQueueWithMax()
		case "enqueue":
			q.Enqueue(o.Arg)
		case "dequeue":
			result := q.Dequeue()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.Arg, result)
			}
		case "max":
			result := q.Max()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.Arg, result)
			}
		}
	}

	return nil
}

type QueueWithMaxOp struct {
	Op  string
	Arg int
}

type queueWithMaxDecoder struct {
	Value []*QueueWithMaxOp
}

func (o *queueWithMaxDecoder) DecodeRecord(record string) error {
	allData := make([][3]interface{}, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([]*QueueWithMaxOp, len(allData))
	for i := 0; i < len(allData); i++ {
		result[i] = &QueueWithMaxOp{
			Op:  allData[i][0].(string),
			Arg: int(allData[i][1].(float64)),
		}
	}

	o.Value = result
	return nil
}
