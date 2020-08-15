package tree_connect_leaves_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/tree_connect_leaves"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestCreateListOfLeaves(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "tree_connect_leaves.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BinaryTreeNodeDecoder
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			result, err := createListOfLeavesWrapper(tc.Tree.Value)
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

func createListOfLeavesWrapper(tree *tree.BinaryTreeNode) ([]int, error) {
	result := CreateListOfLeaves(tree)

	for i, n := range result {
		if n == nil {
			return nil, fmt.Errorf("result contains a nil node at index %d", i)
		}
	}

	extractedRes := make([]int, len(result))
	for i, n := range result {
		extractedRes[i] = n.Data
	}

	return extractedRes, nil
}
