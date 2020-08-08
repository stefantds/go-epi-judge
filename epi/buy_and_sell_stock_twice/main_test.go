package buy_and_sell_stock_twice_test

import (
	"os"
	"testing"

	"github.com/stefantds/go-epi-judge/config"
)

var cfg *config.Config

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.Parse()
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}
