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

// DeepCopyDoubleLinked makes a depp copy of a doubly-linked list.
// It handles linear as well as branching lists.
func DeepCopyDoubleLinked(l *DoublyListNode) *DoublyListNode {
	if l == nil {
		return nil
	}

	nodes := make([]*DoublyListNode, 0)
	copiedNodes := make(map[*DoublyListNode]*DoublyListNode)
	nodes = append(nodes, l)

	for currentIdx := 0; currentIdx < len(nodes); currentIdx++ {
		node := nodes[currentIdx]
		if _, ok := copiedNodes[node]; !ok {

			newNode := DoublyListNode{
				Data: node.Data,
			}

			copiedNodes[node] = &newNode

			if node.Prev != nil {
				nodes = append(nodes, node.Prev)
			}

			if node.Next != nil {
				nodes = append(nodes, node.Next)
			}
		}
	}

	for currentIdx := 0; currentIdx < len(nodes); currentIdx++ {
		node := nodes[currentIdx]
		copy := copiedNodes[node]

		if node.Prev != nil {
			copy.Prev = copiedNodes[node.Prev]
		}

		if node.Next != nil {
			copy.Next = copiedNodes[node.Next]
		}
	}

	return copiedNodes[l]
}

// DoublyListNodeFromSlice transforms an array of values into a doubly linked list.
// The result is non-cyclic and without any branches.
func DoublyListNodeFromSlice(allData []int) *DoublyListNode {
	dummyHead := &DoublyListNode{}
	current := dummyHead

	for _, d := range allData {
		current.Next = &DoublyListNode{
			Data: d,
		}
		current.Next.Prev = current
		current = current.Next
	}

	if dummyHead.Next != nil {
		dummyHead.Next.Prev = nil
	}

	return dummyHead.Next
}

// DoublyListNodeToSlice transforms a doubly linked list into a slice of values.
// It only handles non-cyclic lists. If a cycle is detected, it will panic.
// It only guarantees to include all the values if there are no branches and the given
// node pointer points to the beginning of the list.
func DoublyListNodeToSlice(l *DoublyListNode) []int {
	result := make([]int, 0)
	seenNodes := make(map[*DoublyListNode]bool)
	for cursor := l; cursor != nil; cursor = cursor.Next {
		if ok, _ := seenNodes[cursor]; ok {
			panic("cycle detected")
		}
		seenNodes[cursor] = true
		result = append(result, cursor.Data.(int))
	}
	return result
}
