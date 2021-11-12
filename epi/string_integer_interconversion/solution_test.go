package string_integer_interconversion_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/string_integer_interconversion"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

var solutions = []Solution{
	StringIntegerConverter{},
}

type Solution interface {
	IntToString(x int) string
	StringToInt(s string) int
}

func TestStringIntegerInterconversion(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "string_integer_interconversion.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		IntValue    int
		StringValue string
		Details     string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.IntValue,
			&tc.StringValue,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetTypeName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := stringIntegerInterconversionWrapper(s, tc.IntValue, tc.StringValue); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func stringIntegerInterconversionWrapper(sol Solution, x int, s string) error {
	stringResult := sol.IntToString(x)
	if y, err := strconv.Atoi(stringResult); err != nil || y != x {
		return fmt.Errorf("int to string conversion failed: got: '%s', want: '%s'", stringResult, s)
	}

	if intResult := sol.StringToInt(s); intResult != x {
		return fmt.Errorf("string to int conversion failed: got: %d, want %d", intResult, x)
	}

	return nil
}
