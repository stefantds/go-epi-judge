package tree

type BSTNode struct {
	Data        int
	Left, Right *BSTNode
}

func (b *BSTNode) String() string {
	s, err := binaryTreeToString(b)
	if err != nil {
		panic(err)
	}
	return s
}

func (b BSTNode) GetData() int {
	return b.Data
}

func (b BSTNode) GetLeft() TreeLike {
	return b.Left
}

func (b BSTNode) GetRight() TreeLike {
	return b.Right
}

// DeepCopyBSTNode returns a deep copy of a given BSTNode
func DeepCopyBSTNode(node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	return DeepCopy(node, func(data interface{}, left, right TreeLike) TreeLike {
		var rightNode, leftNode *BSTNode
		if !isNil(right) {
			rightNode = right.(*BSTNode)
		}
		if !isNil(left) {
			leftNode = left.(*BSTNode)
		}

		return &BSTNode{
			Data:  data.(int),
			Left:  leftNode,
			Right: rightNode,
		}
	}).(*BSTNode)
}
