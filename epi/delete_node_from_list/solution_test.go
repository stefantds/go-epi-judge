package delete_node_from_list_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/list"
	. "github.com/stefantds/go-epi-judge/epi/delete_node_from_list"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*list.Node)

var solutions = []solutionFunc{
	DeletionFromList,
}

func TestDeletionFromList(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "delete_node_from_list.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		List           list.NodeDecoder
		NodeIdx        int
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
			&tc.List,
			&tc.NodeIdx,
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
				result := deletionFromListWrapper(s, tc.List.Value, tc.NodeIdx)
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

func deletionFromListWrapper(solution solutionFunc, head *list.Node, nodeIdx int) *list.Node {
	nodeToDelete := head

	if nodeToDelete == nil {
		panic("list is empty")
	}
	for i := nodeIdx; i > 0; i-- {
		if nodeToDelete.Next == nil {
			panic("can't delete last node")
		}

		nodeToDelete = nodeToDelete.Next
	}

	solution(nodeToDelete)
	return head
}
