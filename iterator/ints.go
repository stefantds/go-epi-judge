package iterator

type Ints []int

func (i Ints) Len() int {
	return len(i)
}

func (i Ints) Get(idx int) interface{} {
	return i[idx]
}
