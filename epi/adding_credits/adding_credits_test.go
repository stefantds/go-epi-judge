package adding_credits_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/adding_credits"
)

func TestClientsCreditsInfo(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "adding_credits.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Operations operationsDecoder
		Details    string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
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

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := clientsCreditsInfoTester(tc.Operations.Value); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func clientsCreditsInfoTester(operations []*Operation) error {
	cr := new(ClientsCreditsInfo)

	for opIdx, o := range operations {
		switch o.Op {
		case "ClientsCreditsInfo":
		case "remove":
			var result int
			if cr.Remove(o.SArg) {
				result = 1
			} else {
				result = 0
			}
			if result != o.IArg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.IArg, result)
			}
		case "insert":
			cr.Insert(o.SArg, o.IArg)
		case "add_all":
			cr.AddAll(o.IArg)
		case "lookup":
			result := cr.Lookup(o.SArg)
			if result != o.IArg {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Op, o.IArg, result)
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

func (o *operationsDecoder) DecodeRecord(record string) error {
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
