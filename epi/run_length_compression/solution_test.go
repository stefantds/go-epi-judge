package run_length_compression_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/run_length_compression"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

var solutions = []Solution{
	&RLCompression{},
}

type Solution interface {
	Decoding(s string) string
	Encoding(s string) string
}

func TestRunLengthEncoding(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "run_length_compression.tsv")
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

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetTypeName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := runLengthEncodingTester(s, tc.Encoded, tc.Decoded); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func runLengthEncodingTester(solution Solution, encoded, decoded string) error {
	decodedResult := solution.Decoding(encoded)
	if decodedResult != decoded {
		return fmt.Errorf("decoding failed: got %s, want %s", decodedResult, decoded)
	}

	encodedResult := solution.Encoding(decoded)
	if encodedResult != encoded {
		return fmt.Errorf("encoding failed: got %s, want %s", encodedResult, decoded)
	}

	return nil
}
