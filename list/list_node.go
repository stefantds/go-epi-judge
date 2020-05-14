package list

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Data interface{}
	Next *ListNode
}

func (l ListNode) String() string {
	var buf bytes.Buffer
	for current := &l; current != nil; current = current.Next {
		fmt.Fprintf(&buf, "%v -> ", current.Data)
	}
	fmt.Fprintf(&buf, "%v", nil)
	return buf.String()
}
