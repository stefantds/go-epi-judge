package queue_with_max

type Solution interface {
	Enqueue(x int)
	Dequeue() int
	Max() int
}

type QueueWithMax struct {
	// TODO - Add your code here
}

func NewQueueWithMax() Solution {
	// TODO - Add your code here
	return &QueueWithMax{}
}

func (q *QueueWithMax) Enqueue(x int) {
	// TODO - Add your code here
}

func (q *QueueWithMax) Dequeue() int {
	// TODO - Add your code here
	return 0
}

func (q *QueueWithMax) Max() int {
	// TODO - Add your code here
	return 0
}
