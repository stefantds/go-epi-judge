package interval_add_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/interval_add"
)

func TestAddInterval(t *testing.T) {
	testFileName := filepath.Join(testConfig.TestDataFolder, "interval_add.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		DisjointIntervals intervalsDecoder
		NewInterval       intervalDecoder
		ExpectedResult    intervalsDecoder
		Details           string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.DisjointIntervals,
			&tc.NewInterval,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := AddInterval(tc.DisjointIntervals.Values, tc.NewInterval.Value)
			if !reflect.DeepEqual(result, tc.ExpectedResult.Values) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", result, tc.ExpectedResult.Values)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

type intervalsDecoder struct {
	Values []Interval
}

func (o *intervalsDecoder) DecodeField(record string) error {
	allData := make([][2]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]Interval, len(allData))
	for i := 0; i < len(allData); i++ {
		values[i].Left = allData[i][0]
		values[i].Right = allData[i][1]
	}

	o.Values = values
	return nil
}

type intervalDecoder struct {
	Value Interval
}

func (i *intervalDecoder) DecodeField(record string) error {
	allData := make([]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	i.Value.Left = allData[0]
	i.Value.Right = allData[1]

	return nil
}
