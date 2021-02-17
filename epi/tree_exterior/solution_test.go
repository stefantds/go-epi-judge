package tree_exterior_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
	. "github.com/stefantds/go-epi-judge/epi/tree_exterior"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*tree.BinaryTreeNode) []*tree.BinaryTreeNode

var solutions = []solutionFunc{
	ExteriorBinaryTree,
}

func TestExteriorBinaryTree(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "tree_exterior.tsv")
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result, err := exteriorBinaryTreeWrapper(s, tc.Tree.Value)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func exteriorBinaryTreeWrapper(solution solutionFunc, tree *tree.BinaryTreeNode) ([]int, error) {
	result := solution(tree)
	return createOutputList(result)
}

func createOutputList(l []*tree.BinaryTreeNode) ([]int, error) {
	output := make([]int, 0)

	for _, t := range l {
		if t == nil {
			return nil, errors.New("result list contains nil")
		}
		output = append(output, t.Data)
	}

	return output, nil
}
