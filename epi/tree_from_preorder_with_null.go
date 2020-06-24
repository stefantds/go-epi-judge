package epi

import (
	"github.com/stefantds/go-epi-judge/tree"
)

// IntOrNull represents an integer value that can be nil.
// The value is null if the `Valid` field is false. Otherwise the value is represented
// by the integer value fro the `Value` field.
type IntOrNull struct {
	Value int
	Valid bool
}

func ReconstructPreorder(preorder []IntOrNull) *tree.BinaryTreeNode {
	// TODO - Add your code here
	return nil
}
