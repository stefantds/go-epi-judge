package list

import (
	"bytes"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func (l Node) String() string {
	var buf bytes.Buffer
	visited := make(map[*Node]bool)

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
