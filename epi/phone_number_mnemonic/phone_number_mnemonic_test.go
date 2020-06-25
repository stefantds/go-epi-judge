package phone_number_mnemonic_test

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/phone_number_mnemonic"
)

func TestPhoneMnemonic(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "phone_number_mnemonic.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		PhoneNumber    string
		ExpectedResult []string
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.PhoneNumber,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := PhoneMnemonic(tc.PhoneNumber)
			if !equal(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func equal(result, expected []string) bool {
	sort.Strings(result)
	sort.Strings(expected)

	return reflect.DeepEqual(result, expected)
}
