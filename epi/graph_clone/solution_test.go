package graph_clone_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/graph_clone"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*GraphVertex) *GraphVertex

var solutions = []solutionFunc{
	CloneGraph,
}

func TestCloneGraph(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "graph_clone.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		NumVertices int
		Edges       [][2]int
		Details     string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.NumVertices,
			&tc.Edges,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := cloneGraphWrapper(s, tc.NumVertices, tc.Edges); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func cloneGraphWrapper(solution solutionFunc, numNodes int, edges [][2]int) error {
	graph := newGraph(numNodes, edges)
	result := solution(&graph[0])

	return checkGraph(result, graph)
}

func checkGraph(node *GraphVertex, graph []GraphVertex) error {
	if node == nil {
		return errors.New("graph was not copied")
	}

	vertexSet := make(map[*GraphVertex]bool)
	q := make([]*GraphVertex, 0)

	q = append(q, node)
	vertexSet[node] = true

	var vertex *GraphVertex
	for len(q) > 0 {
		q, vertex = q[:len(q)-1], q[len(q)-1]

		if vertex.Label >= len(graph) {
			return errors.New("invalid vertex number")
		}

		if vertex == &graph[vertex.Label] {
			return errors.New("a vertex from the original graph found in the copy")
		}

		gotLabels := copyLabels(vertex.Edges)
		wantLabels := copyLabels(graph[vertex.Label].Edges)

		sort.Ints(gotLabels)
		sort.Ints(wantLabels)
		if !reflect.DeepEqual(gotLabels, wantLabels) {
			return fmt.Errorf("edges mismatch for vertex %d: expected %v, got %v",
				vertex.Label, wantLabels, gotLabels)
		}

		for _, e := range vertex.Edges {
			if ok := vertexSet[e]; !ok {
				vertexSet[e] = true
				q = append(q, e)
			}
		}
	}

	return nil
}

func copyLabels(edges []*GraphVertex) []int {
	labels := make([]int, len(edges))
	for i, e := range edges {
		labels[i] = e.Label
	}
	return labels
}

func newGraph(numVertices int, edges [][2]int) []GraphVertex {
	result := make([]GraphVertex, numVertices)

	for i := 0; i < len(result); i++ {
		result[i].Label = i
	}

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]

		if from < 0 || from > numVertices-1 || to < 0 || to > numVertices-1 {
			panic(fmt.Errorf("vertex out of bound: %v", edge))
		}
		result[from].Edges = append(result[from].Edges, &result[to])
	}

	return result
}
