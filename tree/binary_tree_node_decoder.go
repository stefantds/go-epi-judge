package tree

import (
	"strconv"
	"strings"
)

type BinaryTreeNodeDecoder struct {
	Tree *BinaryTreeNode
}

// DecodeRecord builds a BinaryTreeNode from a JSON array of ints
func (d *BinaryTreeNodeDecoder) DecodeRecord(record string) error {
	record = strings.TrimPrefix(record, "[")
	record = strings.TrimSuffix(record, "]")
	allData := strings.Split(record, ",")

	nodes := make([]*BinaryTreeNode, 0, len(allData))

	for _, data := range allData {
		n, err := makeNode(strings.TrimSpace(data))
		if err != nil {
			return err
		}

		nodes = append(nodes, n)
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

	d.Tree = root
	return nil
}

func makeNode(value string) (*BinaryTreeNode, error) {
	const nullValue = "null"
	if value == nullValue {
		return nil, nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &BinaryTreeNode{
		Data: i,
	}, nil
}
