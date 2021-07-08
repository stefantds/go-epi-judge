package absent_value_array

// ResetIterator has a metgod Iterator that returns a new channel
// containing the same stream of values every time.
// It effectively allows to reset the channel and read again from it multiple times.
type ResetIterator interface {
	Iterator() <-chan int32
}

func FindMissingElement(stream ResetIterator) int32 {
	// TODO - Add your code here
	return 0
}
