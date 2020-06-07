package epi_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
	"github.com/stefantds/go-epi-judge/list"
)

func TestInsertAfter(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "insert_in_list.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Node           list.ListNodeDecoder
		NodeIdx        int
		NewNodeData    int
		ExpectedResult list.ListNodeDecoder
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Node,
			&tc.NodeIdx,
			&tc.NewNodeData,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result, err := insertListWrapper(tc.Node.Value, tc.NodeIdx, tc.NewNodeData)
			if err != nil {
				t.Error(err)
			} else if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult.Value, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func insertListWrapper(l *list.ListNode, nodeIdx int, newNodeData int) (*list.ListNode, error) {
	node := l

	for nodeIdx > 1 {
		node = node.Next
		nodeIdx -= 1
	}

	newNode := list.ListNode{
		Data: newNodeData,
	}

	InsertAfter(node, &newNode)

	return l, nil
}
