package lowest_common_ancestor_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/lowest_common_ancestor"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestLCA(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "lowest_common_ancestor.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BinaryTreeNodeDecoder
		Key0           int
		Key1           int
		ExpectedResult int
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
			&tc.Key0,
			&tc.Key1,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			result, err := lcaWrapper(tc.Tree.Value, tc.Key0, tc.Key1)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func lcaWrapper(inputTree *tree.BinaryTreeNode, key0 int, key1 int) (int, error) {
	node0 := tree.MustFindNode(inputTree, key0).(*tree.BinaryTreeNode)
	node1 := tree.MustFindNode(inputTree, key1).(*tree.BinaryTreeNode)

	result := LCA(inputTree, node0, node1)

	if result == nil {
		return 0, errors.New("result can not be nil")
	}

	return result.Data.(int), nil
}
