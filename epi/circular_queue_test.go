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

func TestCircularQueue(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "circular_queue.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Operations queueOpDecoder
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
			err := circularQueueWrapper(tc.Operations.Value)
			if err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

func circularQueueWrapper(operations []*QueueOp) error {
	q := NewQueue(1)

	for opIdx, o := range operations {
		switch o.Op {
		case "Queue":
			q = NewQueue(o.Arg)
		case "enqueue":
			q.Enqueue(o.Arg)
		case "dequeue":
			result := q.Dequeue()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.Arg, result)
			}
		case "size":
			result := q.Size()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.Arg, result)
			}
		}
	}

	return nil
}

type QueueOp struct {
	Op  string
	Arg int
}

func (o QueueOp) String() string {
	return fmt.Sprintf("%s %d", o.Op, o.Arg)
}

type queueOpDecoder struct {
	Value []*QueueOp
}

func (o *queueOpDecoder) DecodeRecord(record string) error {
	allData := make([][2]interface{}, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([]*QueueOp, len(allData))
	for i := 0; i < len(allData); i++ {
		result[i] = &QueueOp{
			Op:  allData[i][0].(string),
			Arg: int(allData[i][1].(float64)),
		}
	}

	o.Value = result
	return nil
}
