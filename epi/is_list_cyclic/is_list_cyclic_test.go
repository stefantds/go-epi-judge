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
	testFileName := filepath.Join(testConfig.TestDataFolder, "is_list_cyclic.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Head     list.ListNodeDecoder
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
			if err := hasCycleWrapper(tc.Head.Value, tc.CycleIdx); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func hasCycleWrapper(head *list.ListNode, cycleIdx int) error {
	cycleLength := 0
	if cycleIdx != -1 {
		var cycleStart *list.ListNode
		cursor := head

		for cursor.Next != nil {
			if cursor.Data.(int) == cycleIdx {
				cycleStart = cursor
			}
			cursor = cursor.Next
			if cycleStart != nil {
				cycleLength++
			}
		}

		if cursor.Data == cycleIdx {
			cycleStart = cursor
		}

		cursor.Next = cycleStart
		cycleLength++
	}

	result := HasCycle(head)

	if cycleIdx == -1 {
		if result != nil {
			return fmt.Errorf("expected no cycle, got %d", result.Data.(int))
		}
	} else {
		if result == nil {
			return errors.New("expected a cycle, got nil")
		}

		cursor := result
		for cursor.Next != result {
			cursor = cursor.Next
			cycleLength--

			if cursor == nil || cycleLength <= 0 {
				return errors.New("returned node does not belong to the cycle or is not the closest node to the head")
			}
		}

		// when cursor.Next == result, the remaining length of the cycle should be 1
		if cycleLength != 1 {
			return errors.New("returned node does not belong to the cycle or is not the closest node to the head")
		}
	}

	return nil
}
