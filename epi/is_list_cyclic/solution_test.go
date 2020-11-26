package is_list_cyclic_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/is_list_cyclic"
	"github.com/stefantds/go-epi-judge/list"
)

func TestHasCycle(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "is_list_cyclic.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Head     list.NodeDecoder
		CycleIdx int
		Details  string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Head,
			&tc.CycleIdx,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			if err := hasCycleWrapper(tc.Head.Value, tc.CycleIdx); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func hasCycleWrapper(head *list.Node, cycleIdx int) error {
	var cycleStart *list.Node
	if cycleIdx != -1 {
		cursor := head

		for cursor.Next != nil {
			if cursor.Data == cycleIdx {
				cycleStart = cursor
			}
			cursor = cursor.Next
		}

		if cursor.Data == cycleIdx {
			cycleStart = cursor
		}

		cursor.Next = cycleStart
	}

	result := HasCycle(head)

	if cycleIdx == -1 {
		if result != nil {
			return fmt.Errorf("expected no cycle, got %d", result.Data)
		}
	} else {
		if result == nil {
			return errors.New("expected a cycle, got nil")
		}

		if result != cycleStart {
			return errors.New("returned node is not the node at the start of the cycle")
		}
	}

	return nil
}
