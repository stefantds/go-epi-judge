package epi

type GraphVertex struct {
	Edges []GraphVertex
}

type Edge struct {
	From int
	To int
}

func IsDeadlocked(graph []GraphVertex)bool {
	// TODO - Add your code here
	return false
}
