package tree

import (
	"strconv"
	"strings"
)

type BinaryTreeNodeDecoder struct {
	Value *BinaryTreeNode
}

// DecodeField builds a BinaryTreeNode from a JSON array of ints
func (d *BinaryTreeNodeDecoder) DecodeField(record string) error {
	record = strings.TrimPrefix(record, "[")
	record = strings.TrimSuffix(record, "]")
	allData := strings.Split(record, ",")

	nodes := make([]*BinaryTreeNode, len(allData))

	for i, data := range allData {
		n, err := makeBinaryTreeNode(strings.TrimSpace(data))
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

func makeBinaryTreeNode(value string) (*BinaryTreeNode, error) {
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
