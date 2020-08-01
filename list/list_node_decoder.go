package list

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ListNodeDecoder struct {
	Value *ListNode
}

// DecodeField builds a list from a JSON array of ints
func (d *ListNodeDecoder) DecodeField(record string) error {
	allData := make([]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	dummyHead := &ListNode{}
	current := dummyHead

	for _, d := range allData {
		current.Next = &ListNode{
			Data: d,
		}
		current = current.Next
	}

	d.Value = dummyHead.Next
	return nil
}
