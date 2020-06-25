package task_pairing_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/task_pairing"
)

func TestOptimumTaskAssignment(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "task_pairing.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		TaskDurations  []int
		ExpectedResult [][2]int
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.TaskDurations,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := OptimumTaskAssignment(tc.TaskDurations)
			expectedResult := decodePairedTasks(tc.ExpectedResult)
			if !reflect.DeepEqual(result, expectedResult) {
				t.Errorf("expected %v, got %v", expectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func decodePairedTasks(pairs [][2]int) []PairedTasks {
	result := make([]PairedTasks, len(pairs))

	for i, n := range pairs {
		result[i].Task1 = n[0]
		result[i].Task2 = n[1]
	}

	return result
}
