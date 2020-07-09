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

	d.Value = DoublyListNodeFromSlice(allData)
	return nil
}
