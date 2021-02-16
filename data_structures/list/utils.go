package list

// ToArray transforms a list into an array of values.
// It only handles non-cyclic lists. If a cycle is detected, it will panic.
func ToArray(l *Node) []int {
	result := make([]int, 0)
	seenNodes := make(map[*Node]bool)
	for cursor := l; cursor != nil; cursor = cursor.Next {
		if ok := seenNodes[cursor]; ok {
			panic("cycle detected")
		}
		seenNodes[cursor] = true
		result = append(result, cursor.Data)
	}
	return result
}

// DeepCopy makes a depp copy of a list.
// It handles both non-cyclic and cyclic lists.
func DeepCopy(l *Node) *Node {
	if l == nil {
		return nil
	}

	copiedNodes := make(map[*Node]*Node)
	dummyHead := &Node{}
	cursorCopy := dummyHead
	for cursor := l; cursor != nil; cursor = cursor.Next {
		if copy, ok := copiedNodes[cursor]; ok {
			cursorCopy.Next = copy

			// cycle detected - break the loop
			break
		}
		newNode := Node{
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
func DeepCopyDoubleLinked(l *DoublyLinkedNode) *DoublyLinkedNode {
	if l == nil {
		return nil
	}

	nodes := make([]*DoublyLinkedNode, 0)
	copiedNodes := make(map[*DoublyLinkedNode]*DoublyLinkedNode)
	nodes = append(nodes, l)

	for currentIdx := 0; currentIdx < len(nodes); currentIdx++ {
		node := nodes[currentIdx]
		if _, ok := copiedNodes[node]; !ok {

			newNode := DoublyLinkedNode{
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

// DoublyLinkedNodeFromSlice transforms an array of values into a doubly linked list.
// The result is non-cyclic and without any branches.
func DoublyLinkedNodeFromSlice(allData []int) *DoublyLinkedNode {
	dummyHead := &DoublyLinkedNode{}
	current := dummyHead

	for _, d := range allData {
		current.Next = &DoublyLinkedNode{
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

// DoublyLinkedNodeToSlice transforms a doubly linked list into a slice of values.
// It only handles non-cyclic lists. If a cycle is detected, it will panic.
// It only guarantees to include all the values if there are no branches and the given
// node pointer points to the beginning of the list.
func DoublyLinkedNodeToSlice(l *DoublyLinkedNode) []int {
	result := make([]int, 0)
	seenNodes := make(map[*DoublyLinkedNode]bool)
	for cursor := l; cursor != nil; cursor = cursor.Next {
		if ok := seenNodes[cursor]; ok {
			panic("cycle detected")
		}
		seenNodes[cursor] = true
		result = append(result, cursor.Data)
	}
	return result
}
