package epi_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/goepijudge/epi"
	"github.com/stefantds/goepijudge/tree"
	"github.com/stefantds/goepijudge/[]*tree"
)

func TestCreateListOfLeaves(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "tree_connect_leaves.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree tree.BinaryTreeNodeDecoder
		ExpectedResult []*tree.BinaryTreeNode
		Details string
	}

	parser, err := csv.NewParser(file, &csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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
			result := CreateListOfLeaves(tc.Tree.Tree)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}
