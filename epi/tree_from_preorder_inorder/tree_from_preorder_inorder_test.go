package tree_from_preorder_inorder_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/tree_from_preorder_inorder"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestBinaryTreeFromPreorderInorder(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "tree_from_preorder_inorder.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Preorder       []int
		Inorder        []int
		ExpectedResult tree.BinaryTreeNodeDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Preorder,
			&tc.Inorder,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := BinaryTreeFromPreorderInorder(tc.Preorder, tc.Inorder)
			if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
				t.Errorf("\nexpected:\n%v\ngot:\n%v", tc.ExpectedResult.Value, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}
