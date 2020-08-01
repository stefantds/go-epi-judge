package tree

import (
	"strconv"
	"strings"
)

type BinaryTreeDecoder struct {
	Value *BinaryTree
}

// DecodeField builds a BinaryTree from a JSON array of ints
func (d *BinaryTreeDecoder) DecodeField(record string) error {
	record = strings.TrimPrefix(record, "[")
	record = strings.TrimSuffix(record, "]")
	allData := strings.Split(record, ",")

	nodes := make([]*BinaryTree, len(allData))

	for i, data := range allData {
		n, err := makeBinaryTree(strings.TrimSpace(data))
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
				if current.Left != nil {
					current.Left.Parent = current
				}
			}
			if childrenIdx < len(nodes) {
				current.Right = nodes[childrenIdx]
				childrenIdx++
				if current.Right != nil {
					current.Right.Parent = current
				}
			}
		}
	}

	d.Value = root
	return nil
}

func makeBinaryTree(value string) (*BinaryTree, error) {
	const nullValue = "null"
	if value == nullValue {
		return nil, nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &BinaryTree{
		Data: i,
	}, nil
}
