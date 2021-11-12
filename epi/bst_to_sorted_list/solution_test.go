package bst_to_sorted_list_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
	. "github.com/stefantds/go-epi-judge/epi/bst_to_sorted_list"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*tree.BSTNode) *tree.BSTNode

var solutions = []solutionFunc{
	BstToDoublyLinkedList,
}

func TestBstToDoublyLinkedList(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "bst_to_sorted_list.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BSTNodeDecoder
		ExpectedResult []int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Tree,
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
				result, err := bstToDoublyLinkedListWrapper(s, tree.DeepCopyBSTNode(tc.Tree.Value))
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", result, tc.ExpectedResult, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func bstToDoublyLinkedListWrapper(solution solutionFunc, t *tree.BSTNode) ([]int, error) {
	list := solution(t)

	if list != nil && list.Left != nil {
		return nil, errors.New("function must return the head of the list. Left link must be nil")
	}

	v := make([]int, 0)
	for list != nil {
		v = append(v, list.Data)
		if list.Right != nil && list.Right.Left != list {
			return nil, errors.New("list is ill-formed")
		}
		list = list.Right
	}

	return v, nil
}
