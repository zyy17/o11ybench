package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := New("testdata/config.yaml")
	if err != nil {
		t.Fatalf("failed to create config: %v", err)
	}

	if err := cfg.Print(); err != nil {
		t.Fatalf("failed to print config: %v", err)
	}
}
