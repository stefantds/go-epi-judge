package even_odd_array_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	utils "github.com/stefantds/go-epi-judge/test_utils"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/even_odd_array"
)

func TestEvenOdd(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "even_odd_array.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		A       []int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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
			if cfg.RunParallelTests {
				t.Parallel()
			}
			if err := evenOddWrapper(tc.A); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func evenOddWrapper(a []int) error {
	result := make([]int, len(a))
	copy(result, a)

	EvenOdd(a)

	if err := utils.AssertAllValuesPresent(a, result); err != nil {
		return err
	}

	inOdd := false

	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			if inOdd {
				return fmt.Errorf("even elements appear in odd part")
			}
		} else {
			inOdd = true
		}
	}

	return nil
}
