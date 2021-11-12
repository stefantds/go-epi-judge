package do_lists_overlap_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	"github.com/stefantds/go-epi-judge/data_structures/list"
	. "github.com/stefantds/go-epi-judge/epi/do_lists_overlap"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(*list.Node, *list.Node) *list.Node

var solutions = []solutionFunc{
	OverlappingLists,
}

func TestOverlappingLists(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "do_lists_overlap.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		L0      list.NodeDecoder
		L1      list.NodeDecoder
		Common  list.NodeDecoder
		Cycle0  int
		Cycle1  int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := overlappingListsWrapper(s, tc.L0.Value, tc.L1.Value, tc.Common.Value, tc.Cycle0, tc.Cycle1); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func overlappingListsWrapper(solution solutionFunc, l0 *list.Node, l1 *list.Node, common *list.Node, cycle0 int, cycle1 int) error {
	if common != nil {
		if l0 == nil {
			l0 = common
		} else {
			it := l0
			for it.Next != nil {
				it = it.Next
			}
			it.Next = common
		}

		if l1 == nil {
			l1 = common
		} else {
			it := l1
			for it.Next != nil {
				it = it.Next
			}
			it.Next = common
		}
	}

	if cycle0 != -1 && l0 != nil {
		last := l0
		for last.Next != nil {
			last = last.Next
		}

		it := l0
		for ; cycle0 > 0; cycle0-- {
			if it == nil {
				panic("invalid input data")
			}
			it = it.Next
		}
		last.Next = it
	}

	if cycle1 != -1 && l1 != nil {
		last := l1
		for last.Next != nil {
			last = last.Next
		}

		it := l1
		for ; cycle1 > 0; cycle1-- {
			if it == nil {
				panic("invalid input data")
			}
			it = it.Next
		}
		last.Next = it
	}

	commonNodes := make(map[int]bool)
	for it := common; it != nil; it = it.Next {
		if _, ok := commonNodes[it.Data]; ok {
			break
		}

		commonNodes[it.Data] = true
	}

	result := solution(l0, l1)

	if len(commonNodes) == 0 {
		if result != nil {
			return fmt.Errorf("expected nil, got %v", result)
		}
	} else {
		if result == nil {
			return errors.New("expected a node, got nil")
		}

		_, ok := commonNodes[result.Data]
		if !ok {
			return errors.New("the returned node is not an acceptable answer")
		}
	}

	return nil
}
