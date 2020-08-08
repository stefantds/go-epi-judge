package is_array_dominated_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/is_array_dominated"
)

func TestValidPlacementExists(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "is_array_dominated.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Team0            teamDecoder
		Team1            teamDecoder
		ExpectedResult01 bool
		ExpectedResult10 bool
		Details          string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Team0,
			&tc.Team1,
			&tc.ExpectedResult01,
			&tc.ExpectedResult10,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			if err := validPlacementExistsWrapper(
				tc.Team0.Value,
				tc.Team1.Value,
				tc.ExpectedResult01,
				tc.ExpectedResult10,
			); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func validPlacementExistsWrapper(team0 Team, team1 Team, expected01 bool, expected10 bool) error {
	result01 := ValidPlacementExists(team0, team1)
	if result01 != expected01 {
		return fmt.Errorf("got %t, want %t", result01, expected01)
	}

	result10 := ValidPlacementExists(team1, team0)
	if result10 != expected10 {
		return fmt.Errorf("got %t, want %t", result10, expected10)
	}

	return nil
}

type teamDecoder struct {
	Value Team
}

func (t *teamDecoder) DecodeField(record string) error {
	allData := make([]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	players := make([]Player, len(allData))
	for i, h := range allData {
		players[i] = Player{h}
	}

	t.Value = Team{players}
	return nil
}
