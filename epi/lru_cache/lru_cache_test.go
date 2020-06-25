package lru_cache_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/lru_cache"
)

func TestLRUCache(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "lru_cache.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Operations lruCacheDecoder
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
			if err := lruCacheTester(tc.Operations.Value); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func lruCacheTester(operations []*LRUCacheOp) error {
	var cache LRUCache
	for opIdx, o := range operations {
		switch o.Code {
		case "LruCache":
			cache = NewLRUCache(o.Arg1)
		case "lookup":
			result := cache.Lookup(o.Arg1)
			if result != o.Arg2 {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Code, o.Arg2, result)
			}
		case "insert":
			cache.Insert(o.Arg1, o.Arg2)
		case "erase":
			var result int
			if cache.Erase(o.Arg1) {
				result = 1
			} else {
				result = 0
			}
			if result != o.Arg2 {
				return fmt.Errorf("mismatch at index %d: operation %s: want %d, have %d", opIdx, o.Code, o.Arg2, result)
			}
		}
	}

	return nil
}

type LRUCacheOp struct {
	Code string
	Arg1 int
	Arg2 int
}

func (o LRUCacheOp) String() string {
	return fmt.Sprintf("%s(%d, %d)", o.Code, o.Arg1, o.Arg2)
}

type lruCacheDecoder struct {
	Value []*LRUCacheOp
}

func (o *lruCacheDecoder) DecodeRecord(record string) error {
	allData := make([][3]interface{}, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	result := make([]*LRUCacheOp, len(allData))
	for i := 0; i < len(allData); i++ {
		result[i] = &LRUCacheOp{
			Code: allData[i][0].(string),
			Arg1: int(allData[i][1].(float64)),
			Arg2: int(allData[i][2].(float64)),
		}
	}

	o.Value = result
	return nil
}
