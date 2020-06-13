package iterator

type Interface interface {
	Len() int
	Get(int) interface{}
}

type Iterator struct {
	Data       Interface
	currentIdx int
}

// HasNext returns true if there are more elements to be returned.
func (i *Iterator) HasNext() bool {
	return i.currentIdx < i.Data.Len()
}

// Next returns the next element of the iterator.
// It must be always preceded by a call to HasNext to avoid panic.
func (i *Iterator) Next() interface{} {
	d := i.Data.Get(i.currentIdx)
	i.currentIdx += 1
	return d
}

func New(values Interface) *Iterator {
	return &Iterator{
		Data:       values,
		currentIdx: 0,
	}
}
