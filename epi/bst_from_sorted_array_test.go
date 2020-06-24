package epi_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stefantds/go-epi-judge/utils"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestBuildMinHeightBSTFromSortedArray(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "bst_from_sorted_array.tsv"
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

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result, err := buildMinHeightBSTFromSortedArrayWrapper(tc.A)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(result, tc.ExpectedHeight) {
				t.Errorf("expected %v, got %v", tc.ExpectedHeight, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func buildMinHeightBSTFromSortedArrayWrapper(a []int) (int, error) {
	result := BuildMinHeightBSTFromSortedArray(a)

	inorder := tree.GenerateInorder(result)

	if err := utils.AssertAllValuesPresent(a, inorder); err != nil {
		return 0, err
	}

	if err := tree.AssertTreeIsBST(result); err != nil {
		return 0, err
	}

	return tree.BinaryTreeHeight(result), nil
}
