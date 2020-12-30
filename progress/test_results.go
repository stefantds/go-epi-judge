package main

import (
	"errors"
	"os"
	"path/filepath"

	lib "github.com/stefantds/go-epi-judge/progress/lib"

	"github.com/rogpeppe/go-internal/lockedfile"
)

type resultChecker struct {
	epiFolder string
}

func NewResultChecker(epiFolder string) *resultChecker {
	return &resultChecker{
		epiFolder: epiFolder,
	}
}

// testsPassed checks that all the tests for a particular problem passed.
// The function relies on the tests result file being available.
func (c *resultChecker) testsPassed(p problem) bool {
	folder := filepath.Join(c.epiFolder, p.Folder)

	progressFile := filepath.Join(folder, lib.ProgressFileName)
	rawResult, err := lockedfile.Read(progressFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if _, err := os.Stat(folder); !os.IsNotExist(err) {
				// the folder exists but not the progress file
				return false
			}
		}
		panic(err)
	}

	result := string(rawResult)[0]
	return result == lib.TestSuccessCharacter
}
