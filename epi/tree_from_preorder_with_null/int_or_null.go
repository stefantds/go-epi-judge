package tree_from_preorder_with_null

// IntOrNull represents an integer value that can be nil.
// The value is null if the `Valid` field is false. Otherwise the value is represented
// by the integer value fro the `Value` field.
type IntOrNull struct {
	Value int
	Valid bool
}
