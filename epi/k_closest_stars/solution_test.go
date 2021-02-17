package k_closest_stars_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/k_closest_stars"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(chan Star, int) []Star

var solutions = []solutionFunc{
	FindClosestKStars,
}

func TestFindClosestKStars(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "k_closest_stars.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Stars          starsDecoder
		K              int
		ExpectedResult []float64
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Stars,
			&tc.K,
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
				result := findClosestKStarsWrapper(s, tc.Stars.Value, tc.K)
				if !equal(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func findClosestKStarsWrapper(solution solutionFunc, stars []Star, k int) []Star {
	starsChan := make(chan Star, len(stars))
	for _, s := range stars {
		starsChan <- s
	}
	close(starsChan)

	return solution(starsChan, k)
}

func equal(result []Star, expected []float64) bool {
	if len(expected) != len(result) {
		return false
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Distance() < result[j].Distance() })

	for i := 0; i < len(result); i++ {
		if result[i].Distance() != expected[i] {
			return false
		}
	}

	return true
}

type starsDecoder struct {
	Value []Star
}

func (d *starsDecoder) DecodeField(record string) error {
	allData := make([][3]float64, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	d.Value = make([]Star, len(allData))
	for i := 0; i < len(allData); i++ {
		d.Value[i] = Star{
			X: allData[i][0],
			Y: allData[i][1],
			Z: allData[i][2],
		}
	}

	return nil
}
