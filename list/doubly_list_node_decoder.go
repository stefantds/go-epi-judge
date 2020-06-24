package list

import (
	"encoding/json"
	"fmt"
	"strings"
)

type DoublyListNodeDecoder struct {
	Value *DoublyListNode
}

// DecodeRecord builds a list from a JSON array of ints
func (d *DoublyListNodeDecoder) DecodeRecord(record string) error {
	allData := make([]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	dummyHead := &DoublyListNode{}
	current := dummyHead

	for _, d := range allData {
		current.Next = &DoublyListNode{
			Data: d,
		}
		current.Next.Prev = current
		current = current.Next
	}

	d.Value = dummyHead.Next
	return nil

	// var current DoublyListNode
	// for i := len(allData) - 1; i >= 0; i-- {
	// 	current = DoublyListNode{
	// 		Data: allData[i],
	// 		Next: &current,
	// 	}

	// 	if current.Next != nil {
	// 		current.Next.Prev = &current
	// 	}
	// }

	// d.Value = &current
	// return nil
}
