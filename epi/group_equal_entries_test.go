package epi_test

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestGroupByAge(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "group_equal_entries.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		People  []Person
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
			if err := groupByAgeWrapper(tc.People); err != nil {
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
