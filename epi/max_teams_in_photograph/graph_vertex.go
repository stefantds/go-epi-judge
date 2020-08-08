package max_teams_in_photograph

type GraphVertex struct {
	Edges       []*GraphVertex
	MaxDistance int
}
