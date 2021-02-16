package delete_from_list_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/list"
	. "github.com/stefantds/go-epi-judge/epi/delete_from_list"
)

func TestDeleteList(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "delete_from_list.tsv")
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			result := deleteListWrapper(tc.List.Value, tc.NodeIdx)
			if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult.Value)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func deleteListWrapper(head *list.Node, nodeIdx int) *list.Node {
	nodeToDelete := head
	var prev *list.Node

	if nodeToDelete == nil {
		panic("list is empty")
	}
	for i := nodeIdx; i > 0; i-- {
		if nodeToDelete.Next == nil {
			panic("can't delete last node")
		}

		prev = nodeToDelete
		nodeToDelete = nodeToDelete.Next
	}

	DeleteList(prev)
	return head
}
