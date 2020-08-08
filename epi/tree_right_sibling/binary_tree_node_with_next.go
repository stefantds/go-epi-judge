package tree_right_sibling

type BinaryTreeNodeWithNext struct {
	Data        interface{}
	Left, Right *BinaryTreeNodeWithNext
	Next        *BinaryTreeNodeWithNext
}
