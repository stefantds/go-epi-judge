package pq

// A IntPriorityQueue implements heap.Interface and holds ints
type IntPriorityQueue []int

func (pq IntPriorityQueue) Len() int { return len(pq) }

func (pq IntPriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }

func (pq IntPriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *IntPriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(int)) }

func (pq *IntPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
