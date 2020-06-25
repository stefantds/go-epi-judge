package alternating_array_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stefantds/go-epi-judge/utils"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/alternating_array"
)

func TestRearrange(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "alternating_array.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A       []int
		Details string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.A,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			Rearrange(tc.A)
			if err := rearrangeWrapper(tc.A); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func rearrangeWrapper(a []int) error {
	result := make([]int, len(a))
	_ = copy(result, a)

	Rearrange(result)

	if err := utils.AssertAllValuesPresent(a, result); err != nil {
		return err
	}

	return checkOrder(result)
}

func checkOrder(a []int) error {
	for i := 0; i < len(a); i++ {
		if (i % 2) != 0 {
			if a[i] < a[i-1] {
				return fmt.Errorf("wrong order found: want: a[%d] <= a[%d]; have: %d > %d", i-1, i, a[i-1], a[i])
			}
			if i < len(a)-1 {
				if a[i] < a[i+1] {
					return fmt.Errorf("wrong order found: want: a[%d] >= a[%d]; have: %d < %d", i, i+1, a[i], a[i+1])
				}
			}
		} else {
			if i > 0 {
				if a[i-1] < a[i] {
					return fmt.Errorf("wrong order found: want: a[%d] >= a[%d]; have: %d < %d", i-1, i, a[i-1], a[i])
				}
			}
			if i < len(a)-1 {
				if a[i+1] < a[i] {
					return fmt.Errorf("wrong order found: want: a[%d] <= a[%d]; have: %d > %d", i, i+1, a[i], a[i+1])
				}
			}
		}
	}

	return nil
}
