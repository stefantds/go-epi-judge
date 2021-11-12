package string_transformability_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/string_transformability"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func(map[string]struct{}, string, string) int

var solutions = []solutionFunc{
	TransformString,
}

func TestTransformString(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "string_transformability.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		D              dictDecoder
		S              string
		T              string
		ExpectedResult int
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{
			D: make(dictDecoder),
		}
		if err := parser.Scan(
			&tc.D,
			&tc.S,
			&tc.T,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				result := TransformString(map[string]struct{}(tc.D), tc.S, tc.T)
				if !reflect.DeepEqual(result, tc.ExpectedResult) {
					t.Errorf("\ngot:\n%v\nwant:\n%v\ntest case:\n%+v\n", result, tc.ExpectedResult, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type dictDecoder map[string]struct{}

func (d *dictDecoder) DecodeField(record string) error {
	allData := make([]string, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	for _, v := range allData {
		(*d)[v] = struct{}{}
	}

	return nil
}
