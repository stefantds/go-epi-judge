package tree

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/stefantds/go-epi-judge/data_structures/stack"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type TreeLike interface {
	GetData() int
	GetLeft() TreeLike
	GetRight() TreeLike
}

func isNil(tree TreeLike) bool {
	return tree == nil || reflect.ValueOf(tree).IsNil()
}

func binaryTreeToString(tree TreeLike) (string, error) {
	var buf bytes.Buffer

	nodes := make([]TreeLike, 0)
	visited := make(map[TreeLike]bool)

	first := true
	nullNodesPending := 0

	fmt.Fprint(&buf, "[")
	nodes = append(nodes, tree)

	for currentIdx := 0; currentIdx < len(nodes); currentIdx++ {
		node := nodes[currentIdx]
		if _, found := visited[node]; found {
			return "", errors.New("detected a cycle in the tree")
		}
		if !reflect.ValueOf(node).IsNil() {
			if first {
				first = false
			} else {
				fmt.Fprint(&buf, ", ")
			}

			for nullNodesPending > 0 {
				fmt.Fprint(&buf, "null, ")
				nullNodesPending--
			}

			fmt.Fprintf(&buf, "%v", node.GetData())

			visited[node] = true
			nodes = append(nodes, node.GetLeft())
			nodes = append(nodes, node.GetRight())
		} else {
			nullNodesPending++
		}
	}

	fmt.Fprint(&buf, "]")
	return buf.String(), nil
}

func FindNode(startNode TreeLike, val int) TreeLike {
	s := make(stack.Stack, 0)
	s = s.Push(startNode)

	var node interface{}
	for len(s) > 0 {
		s, node = s.Pop()

		treeNode := node.(TreeLike)
		if isNil(treeNode) {
			continue
		}

		if treeNode.GetData() == val {
			return treeNode
		}

		s = s.Push(treeNode.GetLeft())
		s = s.Push(treeNode.GetRight())
	}

	return nil
}

func MustFindNode(startNode TreeLike, val int) TreeLike {
	n := FindNode(startNode, val)

	if n == nil {
		panic(fmt.Errorf("didn't find the node with value %d in the tree", val))
	}
	return n
}

func GenerateInorder(tree TreeLike) []int {
	result := make([]int, 0)

	if tree == nil {
		return result
	}

	s := make(stack.Stack, 0)
	s = s.Push(tree)

	initial := true
	var node interface{}

	for !s.IsEmpty() {
		s, node = s.Pop()
		treeNode := node.(TreeLike)

		if initial {
			initial = false
		} else {
			result = append(result, treeNode.GetData())
			treeNode = treeNode.GetRight()
		}

		for !isNil(treeNode) {
			s = s.Push(treeNode)
			treeNode = treeNode.GetLeft()
		}
	}

	return result
}

type TreePath struct {
	prev   *TreePath
	toLeft bool
}

func (t *TreePath) WithLeft() *TreePath {
	return &TreePath{
		prev:   t,
		toLeft: true,
	}
}

func (t *TreePath) WithRight() *TreePath {
	return &TreePath{
		prev:   t,
		toLeft: false,
	}
}

type IntRange struct {
	Low  int
	High int
}

func (r *IntRange) contains(value int) bool {
	return r.Low <= value && value <= r.High
}

func (r *IntRange) limitFromBottom(newLow int) *IntRange {
	if newLow > r.Low {
		return &IntRange{newLow, r.High}
	} else {
		return r
	}
}

func (r *IntRange) limitFromTop(newHigh int) *IntRange {
	if newHigh < r.High {
		return &IntRange{r.Low, newHigh}
	} else {
		return r
	}
}

func (r IntRange) String() string {
	return fmt.Sprintf("range between %d and %d", r.Low, r.High)
}

func AssertTreeIsBST(tree TreeLike) error {
	type treePathIntRange struct {
		Tree  TreeLike
		Path  *TreePath
		Range *IntRange
	}

	s := make(stack.Stack, 0)
	s = s.Push(treePathIntRange{
		Tree: tree,
		Path: &TreePath{},
		Range: &IntRange{
			math.MinInt64,
			math.MaxInt64,
		},
	})

	var n interface{}

	for !s.IsEmpty() {
		s, n = s.Pop()
		node := n.(treePathIntRange)
		if isNil(node.Tree) {
			continue
		}

		value := node.Tree.GetData()

		if !node.Range.contains(value) {
			return fmt.Errorf(
				"binary search tree constraints violation: expected value in %s; got %d",
				node.Range,
				value,
			)
		}
		s = s.Push(treePathIntRange{
			Tree:  node.Tree.GetLeft(),
			Path:  node.Path.WithLeft(),
			Range: node.Range.limitFromTop(value),
		})
		s = s.Push(treePathIntRange{
			Tree:  node.Tree.GetRight(),
			Path:  node.Path.WithRight(),
			Range: node.Range.limitFromBottom(value),
		})
	}

	return nil
}

func BinaryTreeHeight(tree TreeLike) int {
	type treeWithHeight struct {
		Tree   TreeLike
		Height int
	}

	s := make(stack.Stack, 0)
	s = s.Push(treeWithHeight{tree, 1})

	height := 0
	var n interface{}

	for !s.IsEmpty() {
		s, n = s.Pop()
		node := n.(treeWithHeight)

		if isNil(node.Tree) {
			continue
		}

		height = utils.Max(height, node.Height)
		s = s.Push(treeWithHeight{
			Tree:   node.Tree.GetLeft(),
			Height: node.Height + 1,
		})
		s = s.Push(treeWithHeight{
			Tree:   node.Tree.GetRight(),
			Height: node.Height + 1,
		})
	}

	return height
}
