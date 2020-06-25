package spiral_ordering_test

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/stefantds/go-epi-judge/config"
)

var testConfig config.TestConfig

func TestMain(m *testing.M) {
	err := parseTestConfig()
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

func parseTestConfig() error {
	f, err := os.Open("../../config.yml")
	if err != nil {
		return fmt.Errorf("can't find config file: %w", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&testConfig)
	if err != nil {
		return fmt.Errorf("can't parse config file: %w", err)
	}

	return nil
}
