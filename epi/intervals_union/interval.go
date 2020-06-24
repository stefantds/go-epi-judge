package intervals_union

import "fmt"

type Interval struct {
	Left  Endpoint
	Right Endpoint
}

func (i Interval) String() string {
	var open, close string
	switch i.Left.IsClosed {
	case true:
		open = "("
	case false:
		open = "["
	}
	switch i.Right.IsClosed {
	case true:
		close = ")"
	case false:
		close = "]"
	}

	return fmt.Sprintf("%s%d, %d%s", open, i.Left.Val, i.Right.Val, close)
}

type Endpoint struct {
	Val      int
	IsClosed bool
}
