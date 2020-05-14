package tree

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

type TreeLike interface {
	GetData() interface{}
	GetLeft() TreeLike
	GetRight() TreeLike
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
