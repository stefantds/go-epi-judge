package epi

type IntervalWithEnds struct {
	Left  Endpoint
	Right Endpoint
}

type Endpoint struct {
	IsClosed bool
	Val      int
}

type FlatInterval struct {
	LeftVal       int
	LeftIsClosed  bool
	RightVal      int
	RightIsClosed bool
}

func UnionOfIntervals(intervals []IntervalWithEnds) []IntervalWithEnds {
	// TODO - Add your code here
	return nil
}
