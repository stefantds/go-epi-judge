package graph_clone

type GraphVertex struct {
	Label int
	Edges []*GraphVertex
}
