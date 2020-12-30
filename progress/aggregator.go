package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

type statsAggregator interface {
	GetStats(chapterKeys []int)
	Render()
}

type detailsAggregator struct {
	chapters problemMapping
	table    *tablewriter.Table
	checker  *resultChecker
}

func NewDetailsAggregator(chapters problemMapping, checker *resultChecker) *detailsAggregator {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Chapter", "Problem", "Tests"})
	t.SetAutoMergeCellsByColumnIndex([]int{0})
	t.SetRowLine(true)
	return &detailsAggregator{
		table:    t,
		chapters: chapters,
		checker:  checker,
	}
}

func (d *detailsAggregator) GetStats(chapterKeys []int) {
	countTotal := 0
	countPassedTotal := 0

	sort.Ints(chapterKeys)
	for _, chapterNb := range chapterKeys {
		chapter := d.chapters[chapterNb]
		countPassedChapter := 0

		for _, prob := range chapter.Problems {
			if testsOK := d.checker.testsPassed(prob); testsOK {
				countPassedChapter++
				d.table.Rich([]string{chapter.Name, prob.Name, "Passed"}, []tablewriter.Colors{{}, {}, {tablewriter.Normal, tablewriter.FgHiGreenColor}})
			} else {
				d.table.Rich([]string{chapter.Name, prob.Name, "Failed"}, []tablewriter.Colors{{}, {}, {tablewriter.Normal, tablewriter.FgHiRedColor}})
			}
		}

		countTotal += len(chapter.Problems)
		countPassedTotal += countPassedChapter
	}

	if countTotal > 0 {
		d.table.SetFooter([]string{"", "Total", fmt.Sprintf("%d/%d (%d%%)", countPassedTotal, countTotal, countPassedTotal*100/countTotal)})
	}
}

func (d *detailsAggregator) Render() {
	d.table.Render()
}

type chaptersAggregator struct {
	chapters problemMapping
	table    *tablewriter.Table
	checker  *resultChecker
}

func NewChaptersAggregator(chapters problemMapping, checker *resultChecker) *chaptersAggregator {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Chapter", "Tests"})
	t.SetRowLine(true)
	return &chaptersAggregator{
		table:    t,
		chapters: chapters,
		checker:  checker,
	}
}

func (d *chaptersAggregator) GetStats(chapterKeys []int) {
	countTotal := 0
	countPassedTotal := 0

	sort.Ints(chapterKeys)
	for _, chapterNb := range chapterKeys {
		chapter := d.chapters[chapterNb]
		countPassedChapter := 0
		for _, prob := range chapter.Problems {
			if testsOK := d.checker.testsPassed(prob); testsOK {
				countPassedChapter++
			}
		}

		countTotal += len(chapter.Problems)
		countPassedTotal += countPassedChapter
		d.table.Append([]string{chapter.Name, fmt.Sprintf("%d/%d (%d%%)", countPassedChapter, len(chapter.Problems), countPassedChapter*100/len(chapter.Problems))})
	}

	if countTotal > 0 {
		d.table.SetFooter([]string{"Total", fmt.Sprintf("%d/%d (%d%%)", countPassedTotal, countTotal, countPassedTotal*100/countTotal)})
	}
}

func (d *chaptersAggregator) Render() {
	d.table.Render()
}
