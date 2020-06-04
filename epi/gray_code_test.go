package epi_test

import (
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestGrayCode(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "gray_code.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		NumBits int
		Details string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.NumBits,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := grayCodeWrapper(tc.NumBits); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func grayCodeWrapper(numBits int) error {
	result := GrayCode(numBits)
	uniqueEntries := make(map[int]bool)

	expectedSize := 1 << numBits
	if expectedSize != len(result) {
		return fmt.Errorf("length mismatch: want %d, have %d", expectedSize, len(result))
	}

	for i := 1; i < len(result); i++ {
		if !differsByOneBit(result[i-1], result[i]) {
			if result[i-1] == result[i] {
				return fmt.Errorf("two consecutive entries are equal at index %d and %d", i-1, i)
			}
			return fmt.Errorf("two consecutive entries differ by more than 1 at index %d and %d", i-1, i)
		}
		if _, ok := uniqueEntries[result[i]]; ok == true {
			return fmt.Errorf("not all entries are distint: %d is duplicated", result[i])
		}
		uniqueEntries[result[i]] = true
	}

	return nil
}

func differsByOneBit(x, y int) bool {
	bitDiff := x ^ y
	return bitDiff != 0 && bitDiff&(bitDiff-1) == 0
}
