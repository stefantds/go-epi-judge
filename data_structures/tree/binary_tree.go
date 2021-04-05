package tree

type BinaryTree struct {
	Data                int
	Left, Right, Parent *BinaryTree
}

func (b *BinaryTree) String() string {
	s, err := binaryTreeToString(b)
	if err != nil {
		panic(err)
	}
	return s
}

func (b BinaryTree) GetData() int {
	return b.Data
}

func (b BinaryTree) GetLeft() TreeLike {
	return b.Left
}

func (b BinaryTree) GetRight() TreeLike {
	return b.Right
}

// DeepCopyBinaryTree returns a deep copy of a given tree
func DeepCopyBinaryTree(node *BinaryTree) *BinaryTree {
	if node == nil {
		return nil
	}
	copy := DeepCopy(node, func(data interface{}, left, right TreeLike) TreeLike {
		var rightNode, leftNode *BinaryTree
		if !isNil(right) {
			rightNode = right.(*BinaryTree)
		}
		if !isNil(left) {
			leftNode = left.(*BinaryTree)
		}

		return &BinaryTree{
			Data:  data.(int),
			Left:  leftNode,
			Right: rightNode,
		}
	})

	copyTree := copy.(*BinaryTree)

	updateParent(copyTree)

	return copyTree
}

// updateParent goes through a BinaryTree and updates the parent field to
// point to the actual parent value.
func updateParent(t *BinaryTree) {
	if t == nil {
		return
	}
	if t.Left != nil {
		t.Left.Parent = t
		updateParent(t.Left)
	}
	if t.Right != nil {
		t.Right.Parent = t
		updateParent(t.Right)
	}
}
