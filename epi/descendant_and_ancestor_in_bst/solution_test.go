package descendant_and_ancestor_in_bst_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
	. "github.com/stefantds/go-epi-judge/epi/descendant_and_ancestor_in_bst"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*tree.BSTNode, *tree.BSTNode, *tree.BSTNode) bool

var solutions = []solutionFunc{
	PairIncludesAncestorAndDescendantOfM,
}

func TestPairIncludesAncestorAndDescendantOfM(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "descendant_and_ancestor_in_bst.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree               tree.BSTNodeDecoder
		PossibleAncOrDesc0 int
		PossibleAncOrDesc1 int
		Middle             int
		ExpectedResult     bool
		Details            string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Tree,
			&tc.PossibleAncOrDesc0,
			&tc.PossibleAncOrDesc1,
			&tc.Middle,
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
				result := pairIncludesAncestorAndDescendantOfMWrapper(s, tc.Tree.Value, tc.PossibleAncOrDesc0, tc.PossibleAncOrDesc1, tc.Middle)
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

func pairIncludesAncestorAndDescendantOfMWrapper(solution solutionFunc, inputTree *tree.BSTNode, possibleAncOrDesc0 int, possibleAncOrDesc1 int, middle int) bool {
	candidate0 := tree.MustFindNode(inputTree, possibleAncOrDesc0).(*tree.BSTNode)
	candidate1 := tree.MustFindNode(inputTree, possibleAncOrDesc1).(*tree.BSTNode)
	middleNode := tree.MustFindNode(inputTree, middle).(*tree.BSTNode)

	return solution(candidate0, candidate1, middleNode)
}
