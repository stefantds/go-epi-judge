package bst_merge_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/tree"
	. "github.com/stefantds/go-epi-judge/epi/bst_merge"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*tree.BSTNode, *tree.BSTNode) *tree.BSTNode

var solutions = []solutionFunc{
	MergeTwoBsts,
}

func TestMergeTwoBsts(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "bst_merge.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A              tree.BSTNodeDecoder
		B              tree.BSTNodeDecoder
		ExpectedResult tree.BSTNodeDecoder
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
			&tc.B,
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
				result := s(tc.A.Value, tc.B.Value)
				if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult.Value)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}
