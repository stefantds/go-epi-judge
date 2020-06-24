package list

// ToArray transforms a list into an array of values.
// It only handles non-cyclic lists. If a cycle is detected, it will panic.
func ToArray(l *ListNode) []int {
	result := make([]int, 0)
	seenNodes := make(map[*ListNode]bool)
	for cursor := l; cursor != nil; cursor = cursor.Next {
		if ok, _ := seenNodes[cursor]; ok {
			panic("cycle detected")
		}
		seenNodes[cursor] = true
		result = append(result, cursor.Data.(int))
	}
	return result
}

// DeepCopy makes a depp copy of a list.
// It handles both non-cyclic and cyclic lists.
func DeepCopy(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}

	copiedNodes := make(map[*ListNode]*ListNode)
	dummyHead := &ListNode{}
	cursorCopy := dummyHead
	for cursor := l; cursor != nil; cursor = cursor.Next {
		if copy, ok := copiedNodes[cursor]; ok {
			cursorCopy.Next = copy

			// cycle detected - break the loop
			break
		}
		newNode := ListNode{
			Data: cursor.Data,
		}
		cursorCopy.Next = &newNode
		copiedNodes[cursor] = &newNode

		cursorCopy = cursorCopy.Next
	}

	return dummyHead.Next
}
