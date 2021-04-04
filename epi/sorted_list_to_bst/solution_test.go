package sorted_list_to_bst_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/iterator"
	"github.com/stefantds/go-epi-judge/data_structures/list"
	. "github.com/stefantds/go-epi-judge/epi/sorted_list_to_bst"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*list.DoublyLinkedNode, int) *list.DoublyLinkedNode

var solutions = []solutionFunc{
	BuildBSTFromSortedList,
}

func TestBuildBSTFromSortedList(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "sorted_list_to_bst.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		DoublyListL list.DoublyLinkedNodeDecoder
		Details     string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}

		if err := parser.Scan(
			&tc.DoublyListL,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := buildBSTFromSortedListWrapper(s, tc.DoublyListL.Value); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func buildBSTFromSortedListWrapper(solution solutionFunc, dl *list.DoublyLinkedNode) error {
	dl = list.DeepCopyDoubleLinked(dl)
	l := list.DoublyLinkedNodeToSlice(dl)
	result := solution(dl, len(l))

	current := iterator.New(iterator.Ints(l))

	if err := compareIterAndTree(result, current); err != nil {
		return err
	}

	if current.HasNext() {
		return errors.New("too few values in the tree")
	}

	return nil
}

func compareIterAndTree(tree *list.DoublyLinkedNode, it *iterator.Iterator) error {
	if tree == nil {
		return nil
	}

	if err := compareIterAndTree(tree.Prev, it); err != nil {
		return err
	}

	if !it.HasNext() {
		return errors.New("too many values in the tree")
	}

	next := it.Next()

	if next != tree.Data {
		return fmt.Errorf("expected value %d, got %d", next.(int), tree.Data)
	}

	if err := compareIterAndTree(tree.Next, it); err != nil {
		return err
	}

	return nil
}
