package main

import (
	"testing"
)

func TestReadConfigWithAllSettings(t *testing.T) {
	config, err := ReadConfig("fixtures/valid_config.yml")

	if err != nil {
		t.Errorf("Failed to read config fixture: %v", err)
	}

	if config.Secure != false {
		t.Errorf("Failed to set Secure from config.")
	}

	if len(config.Redirects) != 1 {
		t.Errorf("Did not parse Redirects from config correctly: %v", config.Redirects)
	}
}
