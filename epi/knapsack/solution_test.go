package knapsack_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/knapsack"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]Item, int) int

var solutions = []solutionFunc{
	OptimumSubjectToCapacity,
}

func TestOptimumSubjectToCapacity(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "knapsack.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Items          itemsDecoder
		Capacity       int
		ExpectedResult int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Items,
			&tc.Capacity,
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
				result := s(tc.Items.Values, tc.Capacity)
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

type itemsDecoder struct {
	Values []Item
}

func (o *itemsDecoder) DecodeField(record string) error {
	allData := make([][2]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]Item, len(allData))
	for i := 0; i < len(allData); i++ {
		values[i].Weight = allData[i][0]
		values[i].Value = allData[i][1]
	}

	o.Values = values
	return nil
}
