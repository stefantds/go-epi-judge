package deadlock_detection

type GraphVertex struct {
	Edges []*GraphVertex
}
