package euclidean_gcd_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/euclidean_gcd"
)

func TestEuclideanGCD(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "gcd.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		X              int64
		Y              int64
		ExpectedResult int64
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.X,
			&tc.Y,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := EuclideanGCD(tc.X, tc.Y)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}
