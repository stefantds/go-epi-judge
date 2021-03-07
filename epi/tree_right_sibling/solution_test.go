package tree_right_sibling_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/tree_right_sibling"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*BinaryTreeNodeWithNext)

var solutions = []solutionFunc{
	ConstructRightSibling,
}

func TestConstructRightSibling(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "tree_right_sibling.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           string
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
			&tc.Tree,
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
				result := constructRightSiblingWrapper(s, tc.Tree)
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

func constructRightSiblingWrapper(solution solutionFunc, encodedTree string) [][]int {
	treeDecoder := binaryTreeNodeWithNextDecoder{}
	_ = treeDecoder.DecodeField(encodedTree)
	tree := treeDecoder.Value

	solution(tree)

	result := make([][]int, 0)
	levelStart := tree

	for levelStart != nil {
		level := make([]int, 0)
		levelIter := levelStart
		for levelIter != nil {
			level = append(level, levelIter.Data)
			levelIter = levelIter.Next
		}
		result = append(result, level)
		levelStart = levelStart.Left
	}

	return result
}

type binaryTreeNodeWithNextDecoder struct {
	Value *BinaryTreeNodeWithNext
}

func (d *binaryTreeNodeWithNextDecoder) DecodeField(record string) error {
	record = strings.TrimPrefix(record, "[")
	record = strings.TrimSuffix(record, "]")
	allData := strings.Split(record, ",")

	nodes := make([]*BinaryTreeNodeWithNext, len(allData))

	for i, data := range allData {
		n, err := makeBinaryTreeNodeWithNext(strings.TrimSpace(data))
		if err != nil {
			return err
		}

		nodes[i] = n
	}

	root := nodes[0]
	childrenIdx := 1

	for nodeIdx := 0; nodeIdx < len(nodes); nodeIdx++ {
		current := nodes[nodeIdx]
		if current != nil {
			if childrenIdx < len(nodes) {
				current.Left = nodes[childrenIdx]
				childrenIdx++
			}
			if childrenIdx < len(nodes) {
				current.Right = nodes[childrenIdx]
				childrenIdx++
			}
		}
	}

	d.Value = root
	return nil
}

func makeBinaryTreeNodeWithNext(value string) (*BinaryTreeNodeWithNext, error) {
	const nullValue = "null"
	if value == nullValue {
		return nil, nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &BinaryTreeNodeWithNext{
		Data: i,
	}, nil
}
