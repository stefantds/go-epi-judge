package epi

import (
	"github.com/stefantds/go-epi-judge/list"
)

func BuildBSTFromSortedList(l *list.DoublyListNode, length int) *list.DoublyListNode {
	head := l
	result, _ := buildSortedListHelper(0, length, head)
	return result
}

func buildSortedListHelper(start, end int, currentHead *list.DoublyListNode) (result, next *list.DoublyListNode) {
	if start >= end {
		return nil, currentHead
	}

	mid := start + (end-start)/2
	left, head := buildSortedListHelper(start, mid, currentHead)

	curr := head
	head.Prev = left
	head = head.Next

	nextCurr, nextHead := buildSortedListHelper(mid+1, end, head)
	curr.Next = nextCurr

	return curr, nextHead
}
