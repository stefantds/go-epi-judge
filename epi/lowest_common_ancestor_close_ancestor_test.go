package epi_test

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestLCAClose(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "lowest_common_ancestor.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BinaryTreeDecoder
		Key0           int
		Key1           int
		ExpectedResult int
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
			&tc.Key0,
			&tc.Key1,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result, err := lcaCloseWrapper(tc.Tree.Value, tc.Key0, tc.Key1)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

func lcaCloseWrapper(inputTree *tree.BinaryTree, key0 int, key1 int) (int, error) {
	node0 := tree.MustFindNode(inputTree, key0).(*tree.BinaryTree)
	node1 := tree.MustFindNode(inputTree, key1).(*tree.BinaryTree)

	result := LCAClose(node0, node1)

	if result == nil {
		return 0, errors.New("result can not be null")
	}

	return result.Data.(int), nil
}
