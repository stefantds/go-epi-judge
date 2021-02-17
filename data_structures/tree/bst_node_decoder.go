package tree

import (
	"strconv"
	"strings"
)

type BSTNodeDecoder struct {
	Value *BSTNode
}

// DecodeField builds a BSTNode from a JSON array of ints
func (d *BSTNodeDecoder) DecodeField(record string) error {
	record = strings.TrimPrefix(record, "[")
	record = strings.TrimSuffix(record, "]")
	allData := strings.Split(record, ",")

	nodes := make([]*BSTNode, len(allData))

	for i, data := range allData {
		n, err := makeBSTNode(strings.TrimSpace(data))
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

func makeBSTNode(value string) (*BSTNode, error) {
	const nullValue = "null"
	if value == nullValue {
		return nil, nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &BSTNode{
		Data: i,
	}, nil
}
