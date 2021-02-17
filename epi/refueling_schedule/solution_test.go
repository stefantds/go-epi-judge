package refueling_schedule_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/refueling_schedule"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]int, []int) int

var solutions = []solutionFunc{
	FindAmpleCity,
}

func TestFindAmpleCity(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "refueling_schedule.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Gallons   []int
		Distances []int
		Details   string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Gallons,
			&tc.Distances,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := findAmpleCityWrapper(s, tc.Gallons, tc.Distances); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findAmpleCityWrapper(solution solutionFunc, gallons []int, distances []int) error {
	result := solution(gallons, distances)
	numCities := len(gallons)
	tank := 0

	for i := 0; i < numCities; i++ {
		city := (result + i) % numCities
		tank += gallons[city]*MPG - distances[city]

		if tank < 0 {
			return fmt.Errorf("out of gas on city %d", city)
		}
	}

	return nil
}
