package string_integer_interconversion_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/string_integer_interconversion"
)

func TestStringIntegerInterconversion(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "string_integer_interconversion.tsv")
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := stringIntegerInterconversionWrapper(tc.IntValue, tc.StringValue); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func stringIntegerInterconversionWrapper(x int, s string) error {
	stringResult := IntToString(x)
	if y, err := strconv.Atoi(stringResult); err != nil || y != x {
		return fmt.Errorf("int to string conversion failed: got: '%s', want: '%s'", stringResult, s)
	}

	if intResult := StringToInt(s); intResult != x {
		return fmt.Errorf("string to int conversion failed: got: %d, want %d", intResult, x)
	}

	return nil
}
