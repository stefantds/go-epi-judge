package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	showDetails       = flag.Bool("v", false, "show detailed test status.")
	epiFolder         = flag.String("d", "./epi", "path to epi folder.")
	requestedChapters []int
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: ./progress [options] [chapter_nb ...]\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Example: ./progress -v -d ../epi 4 5 8\n")
	}

	if err := parseFlags(); err != nil {
		panic(err)
	}

	var agg statsAggregator
	if *showDetails {
		agg = NewDetailsAggregator(chapters, NewResultChecker(*epiFolder))
	} else {
		agg = NewChaptersAggregator(chapters, NewResultChecker(*epiFolder))
	}

	agg.GetStats(getFilteredChapters())
	agg.Render()
}

// parseFlags parses and validates the arguments passed to the cli tool.
func parseFlags() error {
	flag.Parse()

	requestedChapters = make([]int, 0)
	flagChapters := flag.Args()
	for _, k := range flagChapters {
		n, err := strconv.Atoi(k)
		if err != nil {
			return fmt.Errorf("wrong format for chapter number: %v", k)
		}

		if !isChapterNumberValid(n) {
			return fmt.Errorf("chapter number %v doesn't exist", n)
		}
		requestedChapters = append(requestedChapters, n)
	}

	return nil
}

// getFilteredChapters returns the chapter numbers for which the statistics
// are needed.
func getFilteredChapters() []int {
	if len(requestedChapters) > 0 {
		return requestedChapters
	}

	// show all chapters
	return getAllChapterNumbers()
}
