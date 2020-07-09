package group_equal_entries_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/group_equal_entries"
)

func TestGroupByAge(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "group_equal_entries.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		People  personsDecoder
		Details string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.People,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := groupByAgeWrapper(tc.People.Values); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func groupByAgeWrapper(people []Person) error {
	if len(people) == 0 {
		return nil
	}

	values := make(map[Person]int)
	for _, p := range people {
		values[p] += 1
	}

	GroupByAge(people)

	newValues := make(map[Person]int)
	for _, p := range people {
		newValues[p] += 1
	}

	if !reflect.DeepEqual(values, newValues) {
		return errors.New("entries have changed")
	}

	ages := make(map[int]bool)

	lastAge := people[0].Age

	for _, p := range people {
		if ok, _ := ages[p.Age]; ok {
			return errors.New("entries are not grouped by age")
		}

		if p.Age != lastAge {
			ages[lastAge] = true
			lastAge = p.Age
		}
	}

	return nil
}

type personsDecoder struct {
	Values []Person
}

func (o *personsDecoder) DecodeRecord(record string) error {
	allData := make([][2]interface{}, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	values := make([]Person, len(allData))
	for i := 0; i < len(allData); i++ {
		values[i].Age = int(allData[i][0].(float64))
		values[i].Name = allData[i][1].(string)
	}

	o.Values = values
	return nil
}
