package stack

type Stack []interface{}

func (s Stack) Push(v interface{}) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, interface{}) {
	l := len(s)
	return s[:l-1], s[l-1]
}
