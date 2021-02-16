package stack

type Stack []interface{}

func (s Stack) Push(v interface{}) Stack {
	return append(s, v)
}

// Pop returns the element at the top of the stack and
// the stack without the element a the top.
// Will panic if the stack is empty.
func (s Stack) Pop() (Stack, interface{}) {
	if s.IsEmpty() {
		panic("can't pop from an empty stack")
	}
	l := len(s)
	return s[:l-1], s[l-1]
}

// Peek returns the element at the top of the stack.
// The element retrieved does not get removed from the stack.
// Will panic if the stack is empty.
func (s Stack) Peek() interface{} {
	if s.IsEmpty() {
		panic("can't peek on an empty stack")
	}
	return s[len(s)-1]
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
