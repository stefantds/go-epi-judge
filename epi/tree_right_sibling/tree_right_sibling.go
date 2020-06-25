package tree_right_sibling

type BinaryTreeNodeWithNext struct {
	Data        interface{}
	Left, Right *BinaryTreeNodeWithNext
	Next        *BinaryTreeNodeWithNext
}

func ConstructRightSibling(tree *BinaryTreeNodeWithNext) {
	// TODO - Add your code here
}
