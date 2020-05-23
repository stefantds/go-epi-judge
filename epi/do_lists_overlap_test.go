package epi_test

import (
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	// . "github.com/stefantds/go-epi-judge/epi"
	"github.com/stefantds/go-epi-judge/list"
)

func TestOverlappingLists(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "do_lists_overlap.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		L0      list.ListNodeDecoder
		L1      list.ListNodeDecoder
		Common  list.ListNodeDecoder
		Cycle0  int
		Cycle1  int
		Details string
	}

	parser, err := csv.NewParser(file, &csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.L0,
			&tc.L1,
			&tc.Common,
			&tc.Cycle0,
			&tc.Cycle1,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := checkOverlappingLists(tc.L0.Value, tc.L1.Value, tc.Common.Value, tc.Cycle0, tc.Cycle1); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

func checkOverlappingLists(l0 *list.ListNode, l1 *list.ListNode, common *list.ListNode, cycle0 int, cycle1 int) error {
	// TODO
	return nil
}
