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
	visited := make(map[*ListNode]bool)

	for current := &l; current != nil; current = current.Next {
		_, seen := visited[current]

		if seen {
			fmt.Fprintf(&buf, "(cycle to %v)", current.Data)
			return buf.String()
		}

		visited[current] = true
		fmt.Fprintf(&buf, "%v -> ", current.Data)
	}

	fmt.Fprintf(&buf, "%v", nil)
	return buf.String()
}
