package kth_node_in_tree_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
	. "github.com/stefantds/go-epi-judge/epi/kth_node_in_tree"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*BinaryTreeNode, int) *BinaryTreeNode

var solutions = []solutionFunc{
	FindKthNodeBinaryTree,
}

func TestFindKthNodeBinaryTree(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "kth_node_in_tree.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BinaryTreeDecoder
		K              int
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
			&tc.K,
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
				result, err := findKthNodeBinaryTreeWrapper(s, tc.Tree.Value, tc.K)
				if err != nil {
					t.Fatal(err)
				}
				if result != tc.ExpectedResult {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func convertToTreeWithSize(original *tree.BinaryTree) *BinaryTreeNode {
	if original == nil {
		return nil
	}
	left := convertToTreeWithSize(original.Left)
	right := convertToTreeWithSize(original.Right)

	var lSize, rSize int
	if left != nil {
		lSize = left.Size
	}
	if right != nil {
		rSize = right.Size
	}

	return &BinaryTreeNode{
		Data:  original.Data,
		Left:  left,
		Right: right,
		Size:  1 + lSize + rSize,
	}
}

func findKthNodeBinaryTreeWrapper(solution solutionFunc, t *tree.BinaryTree, k int) (int, error) {
	converted := convertToTreeWithSize(t) // no need to deep copy t as a new tree type is created anyway
	result := solution(converted, k)

	if result == nil {
		return 0, errors.New("expected a result, got nil")
	}

	return result.Data, nil
}
