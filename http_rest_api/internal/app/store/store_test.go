package store_test

import (
	"os"
	"testing"
)

var databaseURL string

// TestMain is called once before all the tests
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_name sslmode=disable"
	}

	// TODO: read the docs for string below
	os.Exit(m.Run()) // Need to exit with correct code
}
