package epi

type GraphVertex struct {
	Label int
	Edges []GraphVertex
}

type Edge struct {
	From int
	To int
}

func CloneGraph(graph GraphVertex)GraphVertex {
	// TODO - Add your code here
	return nil
}
