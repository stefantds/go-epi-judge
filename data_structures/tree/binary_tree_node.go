package tree

type BinaryTreeNode struct {
	Data        int
	Left, Right *BinaryTreeNode
}

func (b *BinaryTreeNode) String() string {
	s, err := binaryTreeToString(b)
	if err != nil {
		panic(err)
	}
	return s
}

func (b BinaryTreeNode) GetData() int {
	return b.Data
}

func (b BinaryTreeNode) GetLeft() TreeLike {
	return b.Left
}

func (b BinaryTreeNode) GetRight() TreeLike {
	return b.Right
}

func DeepCopyBinaryTreeNode(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	return DeepCopy(node, func(data interface{}, left, right TreeLike) TreeLike {
		var rightNode, leftNode *BinaryTreeNode
		if !isNil(right) {
			rightNode = right.(*BinaryTreeNode)
		}
		if !isNil(left) {
			leftNode = left.(*BinaryTreeNode)
		}

		return &BinaryTreeNode{
			Data:  data.(int),
			Left:  leftNode,
			Right: rightNode,
		}
	}).(*BinaryTreeNode)
}
