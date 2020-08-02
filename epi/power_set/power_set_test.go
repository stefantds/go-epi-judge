package power_set_test

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/power_set"
	"github.com/stefantds/go-epi-judge/utils"
)

func TestGeneratePowerSet(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "power_set.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		InputSet       []int
		ExpectedResult [][]int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.InputSet,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := GeneratePowerSet(tc.InputSet)
			if !equal(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func equal(result, expected [][]int) bool {
	for _, l := range expected {
		sort.Ints(l)
	}

	sort.Slice(expected, func(i, j int) bool {
		return utils.LexIntsCompare(expected[i], expected[j])
	})

	for _, l := range result {
		sort.Ints(l)
	}

	sort.Slice(result, func(i, j int) bool {
		return utils.LexIntsCompare(result[i], result[j])
	})

	return reflect.DeepEqual(expected, result)
}
