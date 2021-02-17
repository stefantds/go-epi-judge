package search_maze_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/search_maze"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([][]Color, Coordinate, Coordinate) []Coordinate

var solutions = []solutionFunc{
	SearchMaze,
}

func TestSearchMaze(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "search_maze.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Maze           [][]Color
		S              [2]int
		E              [2]int
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
			&tc.Maze,
			&tc.S,
			&tc.E,
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
				result, err := searchMazeWrapper(s, tc.Maze, decodeCoordinate(tc.S), decodeCoordinate(tc.E))
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func searchMazeWrapper(solution solutionFunc, maze [][]Color, s Coordinate, e Coordinate) (bool, error) {
	mazeCopy := make([][]Color, len(maze))

	for i, row := range maze {
		mazeCopy[i] = make([]Color, len(row))
		copy(mazeCopy[i], row)
	}

	path := solution(mazeCopy, s, e)

	if len(path) == 0 {
		return s == e, nil
	}

	if !reflect.DeepEqual(path[0], s) || !reflect.DeepEqual(path[len(path)-1], e) {
		return false, errors.New("path doesn't lay between start and end points")
	}

	for i := 1; i < len(path); i++ {
		if !pathElementIsFeasible(maze, path[i-1], path[i]) {
			return false, errors.New("path contains invalid segments")
		}
	}

	return true, nil
}

func decodeCoordinate(c [2]int) Coordinate {
	return Coordinate{X: c[0], Y: c[1]}
}

func pathElementIsFeasible(maze [][]Color, prev Coordinate, cur Coordinate) bool {
	if !(0 <= cur.X && cur.X < len(maze) && 0 <= cur.Y &&
		cur.Y < len(maze[cur.X]) && maze[cur.X][cur.Y] == 0) {
		return false
	}
	return cur.X == prev.X+1 && cur.Y == prev.Y ||
		cur.X == prev.X-1 && cur.Y == prev.Y ||
		cur.X == prev.X && cur.Y == prev.Y+1 ||
		cur.X == prev.X && cur.Y == prev.Y-1
}
