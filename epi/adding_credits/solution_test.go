package adding_credits_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/adding_credits"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

var solutions = []Solution{
	&ClientsCreditsInfo{},
}

type Solution interface {
	Insert(clientID string, c int)
	Remove(clientID string) bool
	Lookup(clientID string) int
	AddAll(c int)
	Max() string
}

func TestClientsCreditsInfo(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "adding_credits.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Operations operationsDecoder
		Details    string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Operations,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetTypeName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := clientsCreditsInfoTester(s, tc.Operations.Value); err != nil {
					t.Error(err)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func clientsCreditsInfoTester(sol Solution, operations []*Operation) error {
	for opIdx, o := range operations {
		switch o.Op {
		case "ClientsCreditsInfo":
		case "remove":
			var result int
			if sol.Remove(o.SArg) {
				result = 1
			} else {
				result = 0
			}
			if result != o.IArg {
				return fmt.Errorf("mismatch at index %d: operation %s: got: %v, want: %v", opIdx, o.Op, result, o.IArg)
			}
		case "insert":
			sol.Insert(o.SArg, o.IArg)
		case "add_all":
			sol.AddAll(o.IArg)
		case "lookup":
			result := sol.Lookup(o.SArg)
			if result != o.IArg {
				return fmt.Errorf("mismatch at index %d: operation %s: got: %v, want: %v", opIdx, o.Op, result, o.IArg)
			}
		}
	}

	return nil
}

type Operation struct {
	Op   string
	SArg string
	IArg int
}

func (o Operation) String() string {
	return fmt.Sprintf("%s(%s, %d)", o.Op, o.SArg, o.IArg)
}

type operationsDecoder struct {
	Value []*Operation
}

func (o *operationsDecoder) DecodeField(record string) error {
	allData := make([][3]interface{}, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([]*Operation, len(allData))
	for i := 0; i < len(allData); i++ {
		result[i] = &Operation{
			Op:   allData[i][0].(string),
			SArg: allData[i][1].(string),
			IArg: int(allData[i][2].(float64)),
		}
	}

	o.Value = result
	return nil
}
