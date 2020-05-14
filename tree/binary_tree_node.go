package tree

type BinaryTreeNode struct {
	Data        interface{}
	Left, Right *BinaryTreeNode
}

func (b *BinaryTreeNode) String() string {
	s, err := binaryTreeToString(b)
	if err != nil {
		panic(err)
	}
	return s
}

func (b BinaryTreeNode) GetData() interface{} {
	return b.Data
}

func (b BinaryTreeNode) GetLeft() TreeLike {
	return b.Left
}

func (b BinaryTreeNode) GetRight() TreeLike {
	return b.Right
}
