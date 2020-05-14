package tree

type BSTNode struct {
	Data        interface{}
	Left, Right *BSTNode
}

func (b *BSTNode) String() string {
	s, err := binaryTreeToString(b)
	if err != nil {
		panic(err)
	}
	return s
}

func (b BSTNode) GetData() interface{} {
	return b.Data
}

func (b BSTNode) GetLeft() TreeLike {
	return b.Left
}

func (b BSTNode) GetRight() TreeLike {
	return b.Right
}
