package tree_exterior_test

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/tree_exterior"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestExteriorBinaryTree(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "tree_exterior.tsv"
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

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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
			result, err := exteriorBinaryTreeWrapper(tc.Tree.Value)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func exteriorBinaryTreeWrapper(tree *tree.BinaryTreeNode) ([]int, error) {
	result := ExteriorBinaryTree(tree)
	return createOutputList(result)
}

func createOutputList(l []*tree.BinaryTreeNode) ([]int, error) {
	output := make([]int, 0)

	for _, t := range l {
		if t == nil {
			return nil, errors.New("result list contains nil")
		}
		output = append(output, t.Data.(int))
	}

	return output, nil
}