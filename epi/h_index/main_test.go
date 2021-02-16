package h_index_test

import (
	"log"
	"os"
	"testing"

	progress "github.com/stefantds/go-epi-judge/progress/lib"
	"github.com/stefantds/go-epi-judge/test_utils/config"
)

var cfg *config.Config

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.Parse()
	if err != nil {
		panic(err)
	}

	code := m.Run()

	if err = progress.PersistResult(code); err != nil {
		log.Print(err)
	}
	os.Exit(code)
}
