package testresult

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rogpeppe/go-internal/lockedfile"
)

const (
	ProgressFileName = ".progress"
	fileAccessRights = 0777
)

func PersistResult(testOutCode int) error {
	err := lockedfile.Write(
		ProgressFileName,
		strings.NewReader(strconv.Itoa(testOutCode)),
		fileAccessRights,
	)

	if err != nil {
		return fmt.Errorf("could not write progress file: %w", err)
	}

	return nil
}
