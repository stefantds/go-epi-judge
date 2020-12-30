package primitive_multiply_test

import (
	"log"
	"os"
	"testing"

	"github.com/stefantds/go-epi-judge/config"
	progress "github.com/stefantds/go-epi-judge/progress/lib"
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
