package tree_right_sibling_test

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/tree_right_sibling"
)

func TestConstructRightSibling(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "tree_right_sibling.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           binaryTreeNodeWithNextDecoder
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := constructRightSiblingWrapper(tc.Tree.Value)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("\nexpected:\n%v\ngot:\n%v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func constructRightSiblingWrapper(tree *BinaryTreeNodeWithNext) [][]int {
	ConstructRightSibling(tree)

	result := make([][]int, 0)
	levelStart := tree

	for levelStart != nil {
		level := make([]int, 0)
		levelIter := levelStart
		for levelIter != nil {
			level = append(level, levelIter.Data.(int))
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
