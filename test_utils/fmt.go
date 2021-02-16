package test_utils

import (
	"fmt"
	"strings"
)

// MatrixFmt represents a 2D matrix that can be pretty printed
type MatrixFmt struct {
	Val interface{}
}

func (m MatrixFmt) String() string {
	s := fmt.Sprintf("%v", m.Val)
	s = strings.ReplaceAll(s, "[[", "")
	s = strings.ReplaceAll(s, "]]", "")
	s = strings.ReplaceAll(s, "] [", "\n")
	return s
}
