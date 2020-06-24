package epi_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
	"github.com/stefantds/go-epi-judge/iterator"
	"github.com/stefantds/go-epi-judge/list"
)

func TestBuildBSTFromSortedList(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "sorted_list_to_bst.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		RawL        []int
		DoublyListL list.DoublyListNodeDecoder
		Details     string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.RawL,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		if err := parser.Scan(
			&tc.DoublyListL,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := buildBSTFromSortedListWrapper(tc.RawL, tc.DoublyListL.Value); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func buildBSTFromSortedListWrapper(l []int, dl *list.DoublyListNode) error {
	result := BuildBSTFromSortedList(dl, len(l))

	current := iterator.New(iterator.Ints(l))

	if err := compareIterAndTree(result, current); err != nil {
		return err
	}

	if current.HasNext() {
		return errors.New("too few values in the tree")
	}

	return nil
}

func compareIterAndTree(tree *list.DoublyListNode, it *iterator.Iterator) error {
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
		return fmt.Errorf("expected value %d, got %d", next.(int), tree.Data.(int))
	}

	if err := compareIterAndTree(tree.Next, it); err != nil {
		return err
	}

	return nil
}
