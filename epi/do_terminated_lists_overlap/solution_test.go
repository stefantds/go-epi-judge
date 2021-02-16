package do_terminated_lists_overlap_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/list"
	. "github.com/stefantds/go-epi-judge/epi/do_terminated_lists_overlap"
)

func TestOverlappingNoCycleLists(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "do_terminated_lists_overlap.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		FirstPrefix  list.NodeDecoder
		SecondPrefix list.NodeDecoder
		CommonPart   list.NodeDecoder
		Details      string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.FirstPrefix,
			&tc.SecondPrefix,
			&tc.CommonPart,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if cfg.RunParallelTests {
				t.Parallel()
			}
			if err := overlappingNoCycleListsWrapper(
				tc.FirstPrefix.Value,
				tc.SecondPrefix.Value,
				tc.CommonPart.Value,
			); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func overlappingNoCycleListsWrapper(l0 *list.Node, l1 *list.Node, common *list.Node) error {
	if common != nil {
		if l0 != nil {
			i := l0
			for i.Next != nil {
				i = i.Next
			}
			i.Next = common
		} else {
			l0 = common
		}

		if l1 != nil {
			i := l1
			for i.Next != nil {
				i = i.Next
			}
			i.Next = common
		} else {
			l1 = common
		}
	}

	result := OverlappingNoCycleLists(l0, l1)

	if !reflect.DeepEqual(result, list.DeepCopy(common)) {
		return fmt.Errorf("\ngot:\n%v\nwant:\n%v", result, common)
	}

	return nil
}
