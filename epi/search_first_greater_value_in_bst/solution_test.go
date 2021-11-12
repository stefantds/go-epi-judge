package search_first_greater_value_in_bst_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
	. "github.com/stefantds/go-epi-judge/epi/search_first_greater_value_in_bst"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*tree.BSTNode, int) *tree.BSTNode

var solutions = []solutionFunc{
	FindFirstGreaterThanK,
}

func TestFindFirstGreaterThanK(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "search_first_greater_value_in_bst.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BSTNodeDecoder
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
				result := findFirstGreaterThanKWrapper(s, tc.Tree.Value, tc.K)
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

func findFirstGreaterThanKWrapper(solution solutionFunc, node *tree.BSTNode, k int) int {
	node = tree.DeepCopyBSTNode(node)
	if result := solution(node, k); result != nil {
		return result.Data
	}
	return -1
}
