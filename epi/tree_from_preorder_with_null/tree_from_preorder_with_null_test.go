package tree_from_preorder_with_null_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/tree_from_preorder_with_null"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestReconstructPreorder(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "tree_from_preorder_with_null.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Preorder       perorderDecoder
		ExpectedResult tree.BinaryTreeNodeDecoder
		Details        string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Preorder,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := ReconstructPreorder(tc.Preorder.Value)
			if !reflect.DeepEqual(result, tc.ExpectedResult.Value) {
				t.Errorf("\nexpected:\n%v\ngot:\n%v", tc.ExpectedResult.Value, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type perorderDecoder struct {
	Value []IntOrNull
}

func (o *perorderDecoder) DecodeField(record string) error {
	allData := make([]string, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([]IntOrNull, len(allData))
	for i := 0; i < len(allData); i++ {
		switch allData[i] {
		case "null":
			result[i] = IntOrNull{0, false}
		default:
			intVal, err := strconv.Atoi(allData[i])
			if err != nil {
				panic(fmt.Errorf("could not convert %s to int", allData[i]))
			}
			result[i] = IntOrNull{intVal, true}
		}
	}

	o.Value = result
	return nil
}
