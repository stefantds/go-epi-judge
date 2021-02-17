package deadlock_detection_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/deadlock_detection"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]GraphVertex) bool

var solutions = []solutionFunc{
	IsDeadlocked,
}

func TestIsDeadlocked(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "deadlock_detection.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		NumVertices    int
		Edges          [][2]int
		ExpectedResult bool
		Details        string
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
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result := isDeadlockedWrapper(s, tc.NumVertices, tc.Edges)
				if result != tc.ExpectedResult {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func newGraph(numVertices int, edges [][2]int) []GraphVertex {
	result := make([]GraphVertex, numVertices)

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

func isDeadlockedWrapper(solution solutionFunc, numNodes int, edges [][2]int) bool {
	return solution(newGraph(numNodes, edges))
}
