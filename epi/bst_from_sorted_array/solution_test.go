package bst_from_sorted_array_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
	. "github.com/stefantds/go-epi-judge/epi/bst_from_sorted_array"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]int) *tree.BSTNode

var solutions = []solutionFunc{
	BuildMinHeightBSTFromSortedArray,
}

func TestBuildMinHeightBSTFromSortedArray(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "bst_from_sorted_array.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A              []int
		ExpectedHeight int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.A,
			&tc.ExpectedHeight,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result, err := buildMinHeightBSTFromSortedArrayWrapper(s, tc.A)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(result, tc.ExpectedHeight) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedHeight)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func buildMinHeightBSTFromSortedArrayWrapper(solution solutionFunc, a []int) (int, error) {
	result := solution(a)

	inorder := tree.GenerateInorder(result)

	if err := utils.AssertAllValuesPresent(a, inorder); err != nil {
		return 0, err
	}

	if err := tree.AssertTreeIsBST(result); err != nil {
		return 0, err
	}

	return tree.BinaryTreeHeight(result), nil
}
