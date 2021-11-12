package lowest_common_ancestor_in_bst_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
	. "github.com/stefantds/go-epi-judge/epi/lowest_common_ancestor_in_bst"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*tree.BSTNode, *tree.BSTNode, *tree.BSTNode) *tree.BSTNode

var solutions = []solutionFunc{
	FindLCA,
}

func TestFindLca(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "lowest_common_ancestor_in_bst.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BSTNodeDecoder
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result, err := findLCAWrapper(s, tc.Tree.Value, tc.Key0, tc.Key1)
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

func findLCAWrapper(solution solutionFunc, inputTree *tree.BSTNode, key0 int, key1 int) (int, error) {
	inputTree = tree.DeepCopyBSTNode(inputTree)
	node0 := tree.MustFindNode(inputTree, key0).(*tree.BSTNode)
	node1 := tree.MustFindNode(inputTree, key1).(*tree.BSTNode)

	result := solution(inputTree, node0, node1)

	if result == nil {
		return 0, errors.New("result can not be nil")
	}

	return result.Data, nil
}
