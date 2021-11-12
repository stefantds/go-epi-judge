package group_equal_entries_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/group_equal_entries"
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

type solutionFunc = func([]Person)

var solutions = []solutionFunc{
	GroupByAge,
}

func TestGroupByAge(t *testing.T) {
	testFileName := filepath.Join(cfg.TestDataFolder, "group_equal_entries.tsv")
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		People  personsDecoder
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
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

		for _, s := range solutions {
			t.Run(fmt.Sprintf("Test Case %d %v", i, utils.GetFuncName(s)), func(t *testing.T) {
				if cfg.RunParallelTests {
					t.Parallel()
				}
				if err := groupByAgeWrapper(s, tc.People.Values); err != nil {
					t.Errorf("%v\ntest case:\n%+v\n", err, tc)
				}
			})
		}
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func groupByAgeWrapper(solution solutionFunc, people []Person) error {
	result := make([]Person, len(people))
	_ = copy(result, people)

	values := make(map[Person]int)
	for _, p := range result {
		values[p] += 1
	}

	solution(result)

	newValues := make(map[Person]int)
	for _, p := range result {
		newValues[p] += 1
	}

	if !reflect.DeepEqual(values, newValues) {
		return errors.New("entries have changed")
	}

	ages := make(map[int]bool)

	lastAge := result[0].Age

	for _, p := range result {
		if ok := ages[p.Age]; ok {
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

func (o *personsDecoder) DecodeField(record string) error {
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
