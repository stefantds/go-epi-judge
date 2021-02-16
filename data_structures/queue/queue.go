package queue

type Queue []interface{}

func (q Queue) Enqueue(v interface{}) Queue {
	return append(q, v)
}

// Dequeue returns the next element in the queue and the
// queue without the element.
// Will panic if called on an empty queue.
func (q Queue) Dequeue() (Queue, interface{}) {
	if q.IsEmpty() {
		panic("can't dequeue from an empty queue")
	}
	l := len(q)
	return q[1:l], q[0]
}

// Peek returns the the next element in the queue.
// The element retrieved does not get removed from the queue.
// Will panic if called on an empty queue.
func (q Queue) Peek() interface{} {
	if q.IsEmpty() {
		panic("can't peek on an empty queue")
	}
	return q[0]
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
