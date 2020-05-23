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

func TestDeleteList(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "delete_from_list.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		List           list.ListNodeDecoder
		NodeIdx        int
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
			&tc.List,
			&tc.NodeIdx,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := deleteListWrapper(tc.List.Value, tc.NodeIdx)
			if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult.Value, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

func deleteListWrapper(head *list.ListNode, nodeIdx int) *list.ListNode {
	nodeToDelete := head
	var prev *list.ListNode

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
