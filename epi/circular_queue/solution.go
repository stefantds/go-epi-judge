package circular_queue

type Solution interface {
	Enqueue(x int)
	Dequeue() int
	Size() int
}

type Queue struct {
	// TODO - Add your code here
}

func NewQueue(capacity int) Solution {
	// TODO - Add your code here
	return &Queue{}
}

func (q *Queue) Enqueue(x int) {
	// TODO - Add your code here
}

func (q *Queue) Dequeue() int {
	// TODO - Add your code here
	return 0
}

func (q *Queue) Size() int {
	// TODO - Add your code here
	return 0
}
