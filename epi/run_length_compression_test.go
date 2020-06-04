package epi_test

import (
	"fmt"
	"os"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestRunLengthEncoding(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "run_length_compression.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Encoded string
		Decoded string
		Details string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Encoded,
			&tc.Decoded,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := runLengthEncodingTester(tc.Encoded, tc.Decoded); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func runLengthEncodingTester(encoded, decoded string) error {
	decodedResult := Decoding(encoded)
	if decodedResult != decoded {
		return fmt.Errorf("decoding failed: want %s, have %s", decoded, decodedResult)
	}

	encodedResult := Encoding(decoded)
	if encodedResult != encoded {
		return fmt.Errorf("encoding failed: want %s, have %s", decoded, encodedResult)
	}

	return nil
}
