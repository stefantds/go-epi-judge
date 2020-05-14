package epi

type Interval struct {
	Left Endpoint
	Right Endpoint
}

type Endpoint struct {
	IsClosed bool
	Val int
}

type FlatInterval struct {
	LeftVal int
	LeftIsClosed bool
	RightVal int
	RightIsClosed bool
}

func UnionOfIntervals(intervals []Interval)[]Interval {
	// TODO - Add your code here
	return nil
}
