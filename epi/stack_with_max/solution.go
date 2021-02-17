package stack_with_max

type Solution interface {
	Push(x int)
	Pop() int
	Max() int
	Empty() bool
}

type StackWithMax struct {
	// TODO - Add your code here
}

func NewStackWithMax() Solution {
	// TODO - Add your code here
	return &StackWithMax{}
}

func (q *StackWithMax) Push(x int) {
	// TODO - Add your code here
}

func (q *StackWithMax) Pop() int {
	// TODO - Add your code here
	return 0
}

func (q *StackWithMax) Max() int {
	// TODO - Add your code here
	return 0
}

func (q *StackWithMax) Empty() bool {
	// TODO - Add your code here
	return false
}
