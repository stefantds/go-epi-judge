package epi_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
)

func TestCopyPostingsList(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "copy_posting_list.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		L       postingListNodeDecoder
		LCopy   postingListNodeDecoder
		Details string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.L,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		// for simplicity get a deep copy of L by decoding the JSON again
		if err := parser.Scan(
			&tc.LCopy,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := copyPostingsListWrapper(tc.L.Value, tc.LCopy.Value); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

func copyPostingsListWrapper(l *PostingListNode, copyL *PostingListNode) error {
	result := CopyPostingsList(l)
	return checkPostingListsEqual(copyL, result)
}

func checkPostingListsEqual(orig *PostingListNode, copy *PostingListNode) error {
	nodeMap := make(map[*PostingListNode]*PostingListNode)
	oIter := orig
	cIter := copy
	idx := 0

	for oIter != nil {
		switch {
		case cIter == nil:
			return fmt.Errorf("copied list has fewer nodes than the original: want %v, have %v", orig, copy)
		case oIter.Order != cIter.Order:
			return fmt.Errorf("elements mismatch at index %d: want %v, have %v", idx, oIter, cIter)
		}
		nodeMap[oIter] = cIter
		oIter = oIter.Next
		cIter = cIter.Next
		idx += 1
	}

	if cIter != nil {
		return fmt.Errorf("copied list has more nodes than the original: want %v, have %v", orig, copy)
	}

	oIter = orig
	cIter = copy
	idx = 0
	for oIter != nil {
		switch {
		case nodeMap[cIter] != nil:
			return fmt.Errorf("copied list contain a node from the original list at index %d", idx)
		case oIter.Jump != nil && nodeMap[oIter.Jump] != cIter.Jump:
			return fmt.Errorf("jump link points to a different node in the copied list at index %d", idx)
		case oIter.Jump == nil && cIter.Jump != nil:
			return fmt.Errorf("jump link points to a different node in the copied list at index %d", idx)
		}

		oIter = oIter.Next
		cIter = cIter.Next
		idx += 1
	}

	return nil
}

type postingListNodeDecoder struct {
	Value *PostingListNode
}

func (d *postingListNodeDecoder) DecodeRecord(record string) error {
	allData := make([][2]int, 0)
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&allData); err != nil {
		return fmt.Errorf("could not parse %s as JSON array: %w", record, err)
	}

	orderMap := make(map[int]*PostingListNode)
	var head *PostingListNode

	for i := len(allData) - 1; i >= 0; i-- {
		head = &PostingListNode{
			Order: allData[i][0],
			Next:  head,
			Jump:  nil,
		}
		orderMap[head.Order] = head
	}

	listIter := head
	for _, item := range allData {
		if item[1] != -1 {
			listIter.Jump = orderMap[item[1]]
		}
	}

	d.Value = head
	return nil
}
