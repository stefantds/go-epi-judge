package pivot_list_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/pivot_list"
	"github.com/stefantds/go-epi-judge/list"
)

func TestListPivoting(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "pivot_list.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		L       list.NodeDecoder
		X       int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.L,
			&tc.X,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := listPivotingWrapper(tc.L.Value, tc.X); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func listPivotingWrapper(l *list.Node, x int) error {
	original := list.ToArray(l)
	ListPivoting(l, x)
	pivoted := list.ToArray(l)

	const smaller, equal, greater int = 0, 1, 2

	mode := smaller

	for _, i := range pivoted {
		switch mode {
		case smaller:
			switch {
			case i == x:
				mode = equal
			case i > x:
				mode = greater
			}
		case equal:
			switch {
			case i < x:
				return errors.New("result list is not pivoted")
			case i > x:
				mode = greater
			}
		case greater:
			if i <= x {
				return errors.New("result list is not pivoted")
			}
		}
	}

	sort.Ints(original)
	sort.Ints(pivoted)

	if !reflect.DeepEqual(original, pivoted) {
		return errors.New("result list contains different values")
	}

	return nil
}
