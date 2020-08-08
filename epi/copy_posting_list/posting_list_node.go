package copy_posting_list

import (
	"bytes"
	"fmt"
)

type PostingListNode struct {
	Order      int
	Next, Jump *PostingListNode
}

func (l PostingListNode) String() string {
	var buf bytes.Buffer
	for current := &l; current != nil; current = current.Next {
		fmt.Fprintf(&buf, "%v -> ", current.Order)
	}
	fmt.Fprintf(&buf, "%v", nil)
	return buf.String()
}
