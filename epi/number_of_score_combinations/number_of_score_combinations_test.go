package number_of_score_combinations_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/number_of_score_combinations"
)

func TestNumCombinationsForFinalScore(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "number_of_score_combinations.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		FinalScore           int
		IndividualPlayScores []int
		ExpectedResult       int
		Details              string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.FinalScore,
			&tc.IndividualPlayScores,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := NumCombinationsForFinalScore(tc.FinalScore, tc.IndividualPlayScores)
			if result != tc.ExpectedResult {
				t.Errorf("\nexpected:\n%v\ngot:\n%v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}
