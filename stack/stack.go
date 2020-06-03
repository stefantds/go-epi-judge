package stack

type Stack []interface{}

func (s Stack) Push(v interface{}) Stack {
	return append(s, v)
}

// Peek returns the element at the top of the stack and
// the stack without the element a the top.
func (s Stack) Pop() (Stack, interface{}) {
	l := len(s)
	return s[:l-1], s[l-1]
}

// Peek returns the element at the top of the stack.
// The element retrieved does not get removed from the stack.
func (s Stack) Peek() interface{} {
	return s[len(s)-1]
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
