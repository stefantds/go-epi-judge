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
