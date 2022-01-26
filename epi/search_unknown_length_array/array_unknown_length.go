package search_unknown_length_array

type ArrayUnknownLength interface {
	// Get returns the value at the given index in the array, if the index is valid.
	// In this case the valid flag will be true.
	// In case the index is out of bounds, it returns valid = false and an unspecified value.
	Get(index int) (value int, valid bool)
}
