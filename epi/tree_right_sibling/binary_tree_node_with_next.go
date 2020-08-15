package tree_right_sibling

type BinaryTreeNodeWithNext struct {
	Data        int
	Left, Right *BinaryTreeNodeWithNext
	Next        *BinaryTreeNodeWithNext
}
