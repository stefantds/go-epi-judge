package range_lookup_in_bst_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/range_lookup_in_bst"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestRangeLookupInBst(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "range_lookup_in_bst.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BSTNodeDecoder
		Interval       intervalDecoder
		ExpectedResult []int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Tree,
			&tc.Interval,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := RangeLookupInBst(tc.Tree.Value, tc.Interval.Value)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type intervalDecoder struct {
	Value Interval
}

func (i *intervalDecoder) DecodeField(record string) error {
	var allData [2]int
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	i.Value.Left = allData[0]
	i.Value.Right = allData[1]

	return nil
}
