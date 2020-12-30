package testresult

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rogpeppe/go-internal/lockedfile"
)

const (
	ProgressFileName     = ".progress"
	TestSuccessCharacter = '0' // 0 is the success code for a go test. We use the same for convenience
	fileAccessRights     = 0777
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
