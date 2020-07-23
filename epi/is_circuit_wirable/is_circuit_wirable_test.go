package is_circuit_wirable_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/is_circuit_wirable"
)

func TestIsAnyPlacementFeasible(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "is_circuit_wirable.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		K              int
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
			&tc.K,
			&tc.Edges,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := isAnyPlacementFeasibleWrapper(tc.K, tc.Edges)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func isAnyPlacementFeasibleWrapper(k int, edges [][2]int) bool {
	return IsAnyPlacementFeasible(newGraph(k, edges))
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
