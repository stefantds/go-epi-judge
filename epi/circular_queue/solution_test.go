package circular_queue_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/circular_queue"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(int) Solution

var solutions = []solutionFunc{
	NewQueue,
}

func TestCircularQueue(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "circular_queue.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Operations queueOpDecoder
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
				if err := circularQueueTester(s, tc.Operations.Value); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func circularQueueTester(solution solutionFunc, operations []*QueueOp) error {
	q := solution(1)

	for opIdx, o := range operations {
		switch o.Op {
		case "Queue":
		case "enqueue":
			q.Enqueue(o.Arg)
		case "dequeue":
			result := q.Dequeue()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: got: %v, want: %v", opIdx, o.Op, result, o.Arg)
			}
		case "size":
			result := q.Size()
			if result != o.Arg {
				return fmt.Errorf("mismatch at index %d: operation %s: got: %v, want: %v", opIdx, o.Op, result, o.Arg)
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

func (o *queueOpDecoder) DecodeField(record string) error {
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
