package descendant_and_ancestor_in_bst_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/descendant_and_ancestor_in_bst"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestPairIncludesAncestorAndDescendantOfM(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "descendant_and_ancestor_in_bst.tsv"
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

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := pairIncludesAncestorAndDescendantOfMWrapper(tc.Tree.Value, tc.PossibleAncOrDesc0, tc.PossibleAncOrDesc1, tc.Middle)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func pairIncludesAncestorAndDescendantOfMWrapper(inputTree *tree.BSTNode, possibleAncOrDesc0 int, possibleAncOrDesc1 int, middle int) bool {
	candidate0 := tree.MustFindNode(inputTree, possibleAncOrDesc0).(*tree.BSTNode)
	candidate1 := tree.MustFindNode(inputTree, possibleAncOrDesc1).(*tree.BSTNode)
	middleNode := tree.MustFindNode(inputTree, middle).(*tree.BSTNode)

	return PairIncludesAncestorAndDescendantOfM(candidate0, candidate1, middleNode)
}
