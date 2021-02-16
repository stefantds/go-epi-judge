package list

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NodeDecoder struct {
	Value *Node
}

// DecodeField builds a list from a JSON array of ints
func (d *NodeDecoder) DecodeField(record string) error {
	allData := make([]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	dummyHead := &Node{}
	current := dummyHead

	for _, d := range allData {
		current.Next = &Node{
			Data: d,
		}
		current = current.Next
	}

	d.Value = dummyHead.Next
	return nil
}
