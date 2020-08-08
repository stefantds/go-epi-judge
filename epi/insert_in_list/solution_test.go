package insert_in_list_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/insert_in_list"
	"github.com/stefantds/go-epi-judge/list"
)

func TestInsertAfter(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "insert_in_list.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Node           list.NodeDecoder
		NodeIdx        int
		NewNodeData    int
		ExpectedResult list.NodeDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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
			if cfg.RunParallelTests {
				t.Parallel()
			}
			result, err := insertListWrapper(tc.Node.Value, tc.NodeIdx, tc.NewNodeData)
			if err != nil {
				t.Error(err)
			} else if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult.Value)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func insertListWrapper(l *list.Node, nodeIdx int, newNodeData int) (*list.Node, error) {
	node := l

	for nodeIdx > 1 {
		node = node.Next
		nodeIdx -= 1
	}

	newNode := list.Node{
		Data: newNodeData,
	}

	InsertAfter(node, &newNode)

	return l, nil
}
