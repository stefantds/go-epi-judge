package list_cyclic_right_shift_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/list"
	. "github.com/stefantds/go-epi-judge/epi/list_cyclic_right_shift"
)

func TestCyclicallyRightShiftList(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "list_cyclic_right_shift.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		L              list.NodeDecoder
		K              int
		ExpectedResult list.NodeDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.L,
			&tc.K,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			result := CyclicallyRightShiftList(tc.L.Value, tc.K)
			if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult.Value)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}
