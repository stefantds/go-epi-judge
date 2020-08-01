package task_pairing_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

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
		ExpectedResult pairedTasksDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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
			if !reflect.DeepEqual(result, tc.ExpectedResult.Values) {
				t.Errorf("\nexpected:\n%v\ngot:\n%v", tc.ExpectedResult.Values, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type pairedTasksDecoder struct {
	Values []PairedTasks
}

func (d *pairedTasksDecoder) DecodeField(record string) error {
	allData := make([][2]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([]PairedTasks, len(allData))
	for i, n := range allData {
		result[i].Task1 = n[0]
		result[i].Task2 = n[1]
	}

	d.Values = result
	return nil
}
