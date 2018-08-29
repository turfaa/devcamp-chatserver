package server

import (
	"os"
	"testing"

	"chatserver/pkg/lib/config"
)

// TestMain is useful when we need to init common object used widely in tests
// like config object

func TestMain(m *testing.M) {
	cfg = config.Config{
		// since we are testing
		// you can put any value here
	}
	// this will run all the Tests
	os.Exit(m.Run())
}
