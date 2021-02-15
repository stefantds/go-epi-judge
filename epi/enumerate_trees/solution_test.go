package enumerate_trees_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/enumerate_trees"
	"github.com/stefantds/go-epi-judge/tree"
	"github.com/stefantds/go-epi-judge/utils"
)

func TestGenerateAllBinaryTrees(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "enumerate_trees.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		NumNodes       int
		ExpectedResult [][]int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.NumNodes,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			result := generateAllBinaryTreesWrapper(tc.NumNodes)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func generateAllBinaryTreesWrapper(numNodes int) [][]int {
	result := GenerateAllBinaryTrees(numNodes)

	serialized := make([][]int, len(result))

	for i, x := range result {
		serialized[i] = serializeTree(x)
	}

	sort.Slice(serialized, func(i, j int) bool {
		return utils.LexIntsCompare(serialized[i], serialized[j])
	})

	return serialized
}

func serializeTree(t *tree.BinaryTreeNode) []int {
	s := make(utils.Stack, 0)
	s = s.Push(t)

	result := make([]int, 0)
	var n interface{}

	for !s.IsEmpty() {
		s, n = s.Pop()
		p := n.(*tree.BinaryTreeNode)

		if p == nil {
			result = append(result, 0)
		} else {
			result = append(result, 1)
		}

		if p != nil {
			s = s.Push(p.Left)
			s = s.Push(p.Right)
		}
	}

	return result
}
