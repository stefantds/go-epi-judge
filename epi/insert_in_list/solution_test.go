package insert_in_list_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/list"
	. "github.com/stefantds/go-epi-judge/epi/insert_in_list"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*list.Node, *list.Node)

var solutions = []solutionFunc{
	InsertAfter,
}

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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result := insertListWrapper(s, tc.Node.Value, tc.NodeIdx, tc.NewNodeData)
				if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult.Value)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func insertListWrapper(solution solutionFunc, l *list.Node, nodeIdx int, newNodeData int) *list.Node {
	node := l

	for nodeIdx > 1 {
		node = node.Next
		nodeIdx -= 1
	}

	newNode := list.Node{
		Data: newNodeData,
	}

	solution(node, &newNode)

	return l
}
